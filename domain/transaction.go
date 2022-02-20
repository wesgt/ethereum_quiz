package domain

type Transaction struct {
	Hash        string `json:"hash" gorm:"primary_key"`
	BlockHash   string `json:"blockHash"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       uint64 `json:"value"`
	Nonce       uint64 `json:"nonce"`
	Logs        []Log  `gorm:"foreignKey:TxHash"`
	BlockNumber uint64 `json:"-"`
}

type TransactionRepository interface {
	GetByID(hash string) (*Transaction, error)
}

type TransactionUsecase interface {
	GetByID(hash string) (*Transaction, error)
}
