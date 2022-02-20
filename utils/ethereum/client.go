package ethereum

import (
	"context"
	"math/big"

	inTypes "example.com/portto/utils/ethereum/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	rpcClient *rpc.Client
	EthClient *ethclient.Client
}

func Connect(host string) (*Client, error) {
	rpcClient, err := rpc.Dial(host)
	if err != nil {
		return nil, err
	}
	ethClient := ethclient.NewClient(rpcClient)
	return &Client{rpcClient, ethClient}, nil
}

// func (ec *Client) GetBlockNumber(ctx context.Context) (*big.Int, error) {
// 	var result hexutil.Big
// 	err := ec.rpcClient.CallContext(ctx, &result, "eth_blockNumber")
// 	return (*big.Int)(&result), err
// }

func (ec *Client) GetBlockNumber(ctx context.Context) (uint64, error) {
	result, err := ec.EthClient.BlockNumber(ctx)
	return result, err
}

func (ec *Client) GetBlockByNumber(ctx context.Context, number *big.Int) (*inTypes.Block, error) {
	var block *inTypes.Block
	err := ec.rpcClient.CallContext(ctx, &block, "eth_getBlockByNumber", toBlockNumArg(number), false)
	return block, err
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	pending := big.NewInt(-1)
	if number.Cmp(pending) == 0 {
		return "pending"
	}
	return hexutil.EncodeBig(number)
}

func (ec *Client) GetTransactionByHash(ctx context.Context, hash common.Hash) (*inTypes.Transaction, error) {
	var json *inTypes.Transaction
	err := ec.rpcClient.CallContext(ctx, &json, "eth_getTransactionByHash", hash)
	return json, err
}

func (ec *Client) GetTransactionReceipt(ctx context.Context, hash common.Hash) (*inTypes.Receipt, error) {
	var json *inTypes.Receipt
	err := ec.rpcClient.CallContext(ctx, &json, "eth_getTransactionReceipt", hash)
	return json, err
}

// --- for test
func (ec *Client) GetHeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	header, err := ec.EthClient.HeaderByNumber(ctx, number)
	return header, err
}
