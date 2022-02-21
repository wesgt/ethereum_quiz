package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"example.com/portto/config"
	"example.com/portto/utils"
	"example.com/portto/utils/ethereum"
	"example.com/portto/utils/logger"
	"go.uber.org/zap"
)

const workerPoolSize = 4
const scanCount = 40

func init() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	if err := logger.InitLogger(&config.Log); err != nil {
		panic(err)
	}

	if err := utils.InitDB(&config.Database); err != nil {
		panic(err)
	}

	zap.S().Info("Set init end...")
}

func main() {
	lastBlockNum := uint64(0)

	for {

		jobs := make(chan uint64, workerPoolSize)

		// Start workers
		wg := sync.WaitGroup{}
		wg.Add(workerPoolSize)
		for i := 0; i < workerPoolSize; i++ {
			go worker(&wg, i, jobs)
		}

		// add job
		// get current blocknumber
		client, err := ethereum.Connect(config.RPC.Endpoint)
		if err != nil {
			fmt.Println(err.Error())
		}
		blockNum, err := client.GetBlockNumber(context.Background())
		if err != nil {
			zap.S().Error(err.Error())
		}

		startBlockNum := blockNum - scanCount
		if lastBlockNum > 0 {
			startBlockNum = lastBlockNum
		}
		zap.S().Infof("Start blockNum: %d\n", startBlockNum)
		zap.S().Infof("Current blockNum: %d\n", blockNum)

		for i := startBlockNum; i <= blockNum; i++ {
			jobs <- i
		}

		close(jobs)
		wg.Wait()
		lastBlockNum = blockNum
		zap.S().Infof("All workers done")

		time.Sleep(10 * time.Second)
	}

}
