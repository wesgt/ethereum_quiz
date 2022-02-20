package router

import (
	"errors"
	"net/http"

	"example.com/portto/domain"
	"example.com/portto/utils"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	tc domain.TransactionUsecase
}

func NewTransactionHandler(e *gin.Engine, tc domain.TransactionUsecase) {

	handler := &TransactionHandler{tc: tc}
	v1 := e.Group("/transaction")
	{
		v1.GET(":txHash", handler.GetByID)
	}
}

func (handler *TransactionHandler) GetByID(c *gin.Context) {

	txHash := c.Param("txHash")

	if transaction, err := handler.tc.GetByID(txHash); err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		}

	} else {
		logs := []utils.Log{}
		for _, tmp := range transaction.Logs {
			logs = append(logs, utils.Log{
				Index: tmp.Index,
				Data:  tmp.Data,
			})
		}
		result := utils.Transaction{
			Hash:  transaction.Hash,
			From:  transaction.From,
			To:    transaction.To,
			Value: transaction.Value,
			Nonce: transaction.Nonce,
			Logs:  logs,
		}
		c.JSON(http.StatusOK, result)
	}
}
