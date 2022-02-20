package mocks

import (
	"example.com/portto/domain"
	"github.com/stretchr/testify/mock"
)

type BlockRepository struct {
	mock.Mock
}

func (repoMock *BlockRepository) Fetch(limit int) ([]domain.Block, error) {
	args := repoMock.Called(limit)

	if args.Get(0) != nil {
		return args.Get(0).([]domain.Block), args.Error(1)
	} else {
		return []domain.Block{}, args.Error(1)
	}
}

func (repoMock *BlockRepository) GetByID(id int) (*domain.Block, error) {
	args := repoMock.Called(id)

	if args.Get(0) != nil {
		return args.Get(0).(*domain.Block), args.Error(1)
	} else {
		return &domain.Block{}, args.Error(1)
	}

}

func (repoMock *BlockRepository) CreateOrUpdate(block *domain.Block) (*domain.Block, error) {
	return nil, nil
}
