package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Block struct {
	ParentHash   common.Hash    `json:"parentHash"      `
	Hash         common.Hash    `json:"hash"            `
	Number       *hexutil.Big   `json:"number"          `
	Time         hexutil.Uint64 `json:"timestamp"       `
	Transactions []*common.Hash `json:"transactions"    `
}
