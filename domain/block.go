package domain

type Block struct {
	Number       uint64        `json:"block_num" gorm:"primary_key"`
	Hash         string        `json:"block_hash"`
	ParentHash   string        `json:"parent_hash"`
	Time         uint64        `json:"block_time"`
	Transactions []Transaction `gorm:"foreignKey:BlockNumber"`
}

type BlockRepository interface {
	Fetch(limit int) ([]Block, error)
	GetByID(number int) (*Block, error)
	CreateOrUpdate(block *Block) (*Block, error)
}

type BlockUsecase interface {
	Fetch(limit int) ([]Block, error)
	GetByID(number int) (*Block, error)
	CreateByBlockNum(blockNum uint64) error
}
