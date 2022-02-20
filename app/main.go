package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	blockRepository "example.com/portto/block/repository"
	blockRouter "example.com/portto/block/router"
	blockUsecase "example.com/portto/block/usecase"
	"example.com/portto/config"
	"example.com/portto/domain"
	transactionRepository "example.com/portto/transaction/repository"
	transactionRouter "example.com/portto/transaction/router"
	transactionUsecase "example.com/portto/transaction/usecase"
	"example.com/portto/utils"
	"example.com/portto/utils/logger"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	if err := logger.InitLogger(&config.Log); err != nil {
		panic(err)
	}

	if err := utils.InitDB(&config.Database); err != nil {
		panic(err)
	}

	zap.S().Info("Set init end...")
}

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger.Logger, true))

	router.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	// add router
	blockRepo := blockRepository.NewBlockRepository(utils.DB())
	blockUsec := blockUsecase.NewBlockUsecase(blockRepo, nil)
	blockRouter.NewBlockHandler(router, blockUsec)
	transactionRepo := transactionRepository.NewTransactionRepository(utils.DB())
	transactionUsec := transactionUsecase.NewTransactionUsecase(transactionRepo, nil)
	transactionRouter.NewTransactionHandler(router, transactionUsec)

	zap.S().Info("Set router end...")
	return router
}

func main() {
	// DB AutoMigrate
	if err := utils.DB().AutoMigrate(
		new(domain.Block),
		new(domain.Transaction),
		new(domain.Log)); err != nil {
		panic("Failed to migrate db, err : " + err.Error())
	}
	zap.S().Info("DB AutoMigrate end...")

	router := setupRouter()

	srv := &http.Server{
		Addr:    config.Server.Address,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			zap.S().Infof("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Fatalf("Server forced to shutdown:", err)
	}

	zap.S().Info("Server exiting")
}
