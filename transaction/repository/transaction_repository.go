package repository

import (
	"errors"

	"example.com/portto/domain"
	"example.com/portto/utils"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (repo *transactionRepository) GetByID(hash string) (*domain.Transaction, error) {
	var transaction domain.Transaction

	if result := repo.db.Preload("Logs").First(&transaction, "hash=?", hash); result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, result.Error
	}

	return &transaction, nil
}
