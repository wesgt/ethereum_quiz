package router

import (
	"errors"
	"net/http"
	"strconv"

	"example.com/portto/domain"
	"example.com/portto/utils"
	"github.com/gin-gonic/gin"
)

type BlockHandler struct {
	bc domain.BlockUsecase
}

func NewBlockHandler(e *gin.Engine, bc domain.BlockUsecase) {

	handler := &BlockHandler{bc: bc}
	v1 := e.Group("/blocks")
	{
		v1.GET("", handler.Fetch)
		v1.GET(":id", handler.GetByID)
	}
}

func (handler *BlockHandler) Fetch(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "1"))

	if blocks, err := handler.bc.Fetch(limit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	} else {
		results := []utils.Block{}
		for _, block := range blocks {
			tmpBlock := utils.Block{
				Number:     block.Number,
				Hash:       block.Hash,
				ParentHash: block.ParentHash,
				Time:       block.Time,
			}
			results = append(results, tmpBlock)
		}
		c.JSON(http.StatusOK, results)
	}
}

func (handler *BlockHandler) GetByID(c *gin.Context) {

	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"mag": utils.ErrNotFound})
	} else {
		if block, err := handler.bc.GetByID(id); err != nil {
			if errors.Is(err, utils.ErrNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			}

		} else {
			transactions := []string{}
			for _, transaction := range block.Transactions {
				transactions = append(transactions, transaction.Hash)
			}
			result := utils.BlockWithTx{
				Transactions: transactions,
				Block: utils.Block{
					Number:     block.Number,
					Hash:       block.Hash,
					ParentHash: block.ParentHash,
					Time:       block.Time,
				},
			}
			c.JSON(http.StatusOK, result)
		}
	}
}
