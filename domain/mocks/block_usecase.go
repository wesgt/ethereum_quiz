package mocks

import (
	"example.com/portto/domain"
	"github.com/stretchr/testify/mock"
)

type BlockUsecase struct {
	mock.Mock
}

func (uCaseMock *BlockUsecase) Fetch(limit int) ([]domain.Block, error) {
	args := uCaseMock.Called(limit)

	if args.Get(0) != nil {
		return args.Get(0).([]domain.Block), args.Error(1)
	} else {
		return []domain.Block{}, args.Error(1)
	}
}

func (uCaseMock *BlockUsecase) GetByID(id int) (*domain.Block, error) {
	args := uCaseMock.Called(id)

	if args.Get(0) != nil {
		return args.Get(0).(*domain.Block), args.Error(1)
	} else {
		return &domain.Block{}, args.Error(1)
	}

}

func (uCaseMock *BlockUsecase) CreateByBlockNum(blockNum uint64) error {
	return nil
}
