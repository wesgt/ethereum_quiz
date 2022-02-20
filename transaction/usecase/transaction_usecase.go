package usecase

import (
	"example.com/portto/domain"
	"example.com/portto/utils/ethereum"
)

type transactionUsecase struct {
	repo   domain.TransactionRepository
	client *ethereum.Client
}

func NewTransactionUsecase(repo domain.TransactionRepository, client *ethereum.Client) domain.TransactionUsecase {
	return &transactionUsecase{
		repo:   repo,
		client: client,
	}
}

func (u *transactionUsecase) GetByID(hash string) (*domain.Transaction, error) {
	if transaction, err := u.repo.GetByID(hash); err != nil {
		return nil, err
	} else {
		return transaction, nil
	}
}
