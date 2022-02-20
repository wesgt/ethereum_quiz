package repository_test

import (
	"regexp"
	"testing"

	"example.com/portto/block/repository"
	"example.com/portto/domain"
	"example.com/portto/utils"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var mockBlock = &domain.Block{
	Number:       16683937,
	Hash:         "0xee1e0edcbc8882dd1c416a3a4c98ed515fbdaaa497a50229d3da9abd8de06862",
	ParentHash:   "0xaddc8550b7c9c2adb5a21f94b1ccfb992d966cee4eed6bb74c1129b8333e8c70",
	Time:         1644651061,
	Transactions: []domain.Transaction{{Hash: "1"}},
}

func TestFetch(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()
	db := utils.NewMockDB(t, sqlDB)

	blockRows := sqlmock.NewRows([]string{"number", "hash", "parent_hash", "time", "transaction_hash"}).
		AddRow(mockBlock.Number, mockBlock.Hash, mockBlock.ParentHash,
			mockBlock.Time, mockBlock.Transactions[0].Hash)

	blockQuery := "SELECT * FROM `blocks` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(blockQuery)).WillReturnRows(blockRows)

	blockRepo := repository.NewBlockRepository(db)
	blocks, err := blockRepo.Fetch(1)
	assert.NoError(t, err)
	assert.NotNil(t, blocks)
}

func TestGetByID(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()
	db := utils.NewMockDB(t, sqlDB)

	blockRows := sqlmock.NewRows([]string{"number", "hash", "parent_hash", "time", "transaction_hash"}).
		AddRow(mockBlock.Number, mockBlock.Hash, mockBlock.ParentHash,
			mockBlock.Time, mockBlock.Transactions[0].Hash)

	txRows := sqlmock.NewRows([]string{"hash", "block_number"}).
		AddRow(mockBlock.Transactions[0].Hash, mockBlock.Number)

	blockQuery := "SELECT * FROM `blocks` WHERE `blocks`.`number` = ? ORDER BY `blocks`.`number` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(blockQuery)).WithArgs(mockBlock.Number).WillReturnRows(blockRows)

	txQuery := "SELECT `hash`,`block_number` FROM `transactions` WHERE `transactions`.`block_number` = ?"
	mock.ExpectQuery(regexp.QuoteMeta(txQuery)).WithArgs(mockBlock.Number).WillReturnRows(txRows)

	blockRepo := repository.NewBlockRepository(db)
	block, err := blockRepo.GetByID(int(mockBlock.Number))
	assert.NoError(t, err)
	assert.NotNil(t, block)
}
