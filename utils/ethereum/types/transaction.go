package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Transaction struct {
	Hash      common.Hash    `json:"hash" `
	BlockHash common.Hash    `json:"blockHash" `
	From      common.Address `json:"from"      `
	To        common.Address `json:"to"        `
	Value     hexutil.Uint64 `json:"value"     `
	Nonce     hexutil.Uint64 `json:"nonce"     `
}

type Receipt struct {
	TxHash common.Hash    `json:"transactionHash" `
	From   common.Address `json:"from"            `
	To     common.Address `json:"to"              `
	Logs   []*Log         `json:"logs"            `
}

type Log struct {
	Data  hexutil.Bytes `json:"data"    `
	Index hexutil.Uint  `json:"logIndex"`
}
