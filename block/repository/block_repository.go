package repository

import (
	"errors"

	"example.com/portto/domain"
	"example.com/portto/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type blockRepository struct {
	db *gorm.DB
}

func NewBlockRepository(db *gorm.DB) domain.BlockRepository {
	return &blockRepository{
		db: db,
	}
}

func (repo *blockRepository) Fetch(limit int) ([]domain.Block, error) {
	var blocks []domain.Block

	if result := repo.db.Limit(limit).Find(&blocks); result.Error != nil {
		return nil, result.Error
	}

	return blocks, nil
}

func (repo *blockRepository) GetByID(id int) (*domain.Block, error) {
	var block domain.Block

	if result := repo.db.Preload("Transactions",
		func(db *gorm.DB) *gorm.DB {
			return db.Select("hash", "block_number")
		}).First(&block, id); result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, result.Error
	}

	return &block, nil
}

func (repo *blockRepository) CreateOrUpdate(block *domain.Block) (*domain.Block, error) {
	// result := repo.db.Debug().Clauses(clause.OnConflict{
	result := repo.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&block)

	if result.Error != nil {
		return nil, result.Error
	}
	return block, nil
}
