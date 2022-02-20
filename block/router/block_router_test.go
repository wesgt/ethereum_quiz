package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	blockRouter "example.com/portto/block/router"
	"example.com/portto/domain"
	"example.com/portto/domain/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockBlock = domain.Block{
	Number:       16683937,
	Hash:         "0xee1e0edcbc8882dd1c416a3a4c98ed515fbdaaa497a50229d3da9abd8de06862",
	ParentHash:   "0xaddc8550b7c9c2adb5a21f94b1ccfb992d966cee4eed6bb74c1129b8333e8c70",
	Time:         1644651061,
	Transactions: []domain.Transaction{{Hash: "1"}},
}

func TestFetch(t *testing.T) {
	router := gin.New()

	mockUCase := new(mocks.BlockUsecase)
	mockListBlock := []domain.Block{mockBlock}
	mockUCase.On("Fetch", mock.AnythingOfType("int")).Return(mockListBlock, nil)
	blockRouter.NewBlockHandler(router, mockUCase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blocks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUCase.AssertExpectations(t)
}
