package repository

import (
	"regexp"
	"testing"

	"example.com/portto/domain"
	"example.com/portto/utils"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var mockTransaction = &domain.Transaction{
	Hash:      "0x15486e9436bea442c8b3269d7e0e4ad2636efe54224ab4cd48243eb81632fa84",
	BlockHash: "",
	From:      "0xb0fE4FB305bc1A2de833fe85C71d129485226ff5",
	To:        "0x3c16efCd766764F1bdd67dDe2F0b4D24AC30aa0A",
	Value:     0,
	Nonce:     696,
	Logs:      []domain.Log{{Index: 14}},
}

func TestGetByID(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()
	db := utils.NewMockDB(t, sqlDB)

	txRows := sqlmock.NewRows([]string{"hash", "block_hash", "from", "to", "value", "nonce", "log_index"}).
		AddRow(mockTransaction.Hash, mockTransaction.BlockHash, mockTransaction.From,
			mockTransaction.To, mockTransaction.Value, mockTransaction.Nonce, mockTransaction.Logs[0].Index)

	logRows := sqlmock.NewRows([]string{"log_index", "tx_hash"}).
		AddRow(mockTransaction.Logs[0].Index, mockTransaction.Hash)

	txQuery := "SELECT * FROM `transactions` WHERE hash=? ORDER BY `transactions`.`hash` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(txQuery)).WithArgs(mockTransaction.Hash).WillReturnRows(txRows)

	logQuery := "SELECT * FROM `logs` WHERE `logs`.`tx_hash` = ?"
	mock.ExpectQuery(regexp.QuoteMeta(logQuery)).WithArgs(mockTransaction.Hash).WillReturnRows(logRows)

	transactionRepo := NewTransactionRepository(db)
	transactions, err := transactionRepo.GetByID(mockTransaction.Hash)
	assert.NoError(t, err)
	assert.NotNil(t, transactions)
}
