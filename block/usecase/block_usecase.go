package usecase

import (
	"context"
	"math/big"

	"example.com/portto/domain"
	"example.com/portto/utils/ethereum"
	"go.uber.org/zap"
)

type blockUsecase struct {
	repo   domain.BlockRepository
	client *ethereum.Client
}

func NewBlockUsecase(repo domain.BlockRepository, client *ethereum.Client) domain.BlockUsecase {
	return &blockUsecase{
		repo:   repo,
		client: client,
	}
}

func (u *blockUsecase) Fetch(limit int) ([]domain.Block, error) {
	if users, err := u.repo.Fetch(limit); err != nil {
		return []domain.Block{}, err
	} else {
		return users, nil
	}
}

func (u *blockUsecase) GetByID(id int) (*domain.Block, error) {
	if block, err := u.repo.GetByID(id); err != nil {
		return nil, err
	} else {
		return block, nil
	}
}

func (u *blockUsecase) CreateByBlockNum(blockNum uint64) error {
	zap.S().Infof("Create blockNum: %d", blockNum)

	etBlock, err := u.client.GetBlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
	if err != nil {
		return err
	}
	// zap.S().Infof("Number: %v", etBlock.Number.ToInt())
	// zap.S().Infof("Hash: %v", etBlock.Hash.String())
	// zap.S().Infof("Time: %v", (uint64)(etBlock.Time))
	// zap.S().Infof("Transactions[0]: %v", *etBlock.Transactions[0])

	// Get Transaction information
	transactions := []domain.Transaction{}

	for _, txHash := range etBlock.Transactions {
		zap.S().Infof("blockNum: %d, Transaction: %v", blockNum, txHash)

		transaction, err := u.client.GetTransactionByHash(context.Background(), *txHash)
		if err != nil {
			continue
		}

		// Get Transaction logs
		receipt, err := u.client.GetTransactionReceipt(context.Background(), *txHash)
		if err != nil {
			return err
		}
		// logsData, _ := json.Marshal(receipt.Logs)
		// fmt.Println(string(logsData))

		logs := []domain.Log{}
		for _, log := range receipt.Logs {
			logs = append(logs, domain.Log{
				Index: uint(log.Index),
				Data:  log.Data.String(),
			})
		}

		transactions = append(transactions, domain.Transaction{
			Hash:      txHash.String(),
			BlockHash: transaction.BlockHash.String(),
			From:      transaction.From.String(),
			To:        transaction.To.String(),
			Value:     uint64(transaction.Value),
			Nonce:     uint64(transaction.Nonce),
			Logs:      logs,
		})
	}

	u.repo.CreateOrUpdate(&domain.Block{
		Number:       etBlock.Number.ToInt().Uint64(),
		Hash:         etBlock.Hash.String(),
		ParentHash:   etBlock.ParentHash.String(),
		Time:         (uint64)(etBlock.Time),
		Transactions: transactions,
	})

	return nil
}
