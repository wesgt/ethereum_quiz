package usecase_test

import (
	"errors"
	"testing"

	"example.com/portto/block/usecase"
	"example.com/portto/domain"
	"example.com/portto/domain/mocks"
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
	mockBlockRepo := new(mocks.BlockRepository)
	mockListBlock := []domain.Block{mockBlock}

	t.Run("success", func(t *testing.T) {
		mockBlockRepo.On(
			"Fetch",
			mock.AnythingOfType("int")).Return(mockListBlock, nil).Once()

		blockUsec := usecase.NewBlockUsecase(mockBlockRepo, nil)
		limit := 1
		list, err := blockUsec.Fetch(limit)

		assert.NoError(t, err)
		assert.Len(t, list, len(mockListBlock))

		mockBlockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockBlockRepo.On(
			"Fetch",
			mock.AnythingOfType("int")).Return(nil, errors.New("Unexpexted Error")).Once()

		blockUsec := usecase.NewBlockUsecase(mockBlockRepo, nil)
		limit := 1
		list, err := blockUsec.Fetch(limit)

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockBlockRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	mockBlockRepo := new(mocks.BlockRepository)

	t.Run("success", func(t *testing.T) {
		mockBlockRepo.On(
			"GetByID",
			mock.AnythingOfType("int")).Return(&mockBlock, nil).Once()

		blockUsec := usecase.NewBlockUsecase(mockBlockRepo, nil)
		id := 16683937
		block, err := blockUsec.GetByID(id)

		assert.NoError(t, err)
		assert.NotNil(t, block)

		mockBlockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockBlockRepo.On(
			"GetByID",
			mock.AnythingOfType("int")).Return(nil, errors.New("Unexpexted Error")).Once()

		blockUsec := usecase.NewBlockUsecase(mockBlockRepo, nil)
		id := 16683937
		block, err := blockUsec.GetByID(id)

		assert.Error(t, err)
		assert.Nil(t, block)
		mockBlockRepo.AssertExpectations(t)
	})
}
