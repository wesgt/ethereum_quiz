package main

import (
	"sync"

	blockRepository "example.com/portto/block/repository"
	blockUsecase "example.com/portto/block/usecase"
	"example.com/portto/config"
	"example.com/portto/utils"
	"example.com/portto/utils/ethereum"
	"go.uber.org/zap"
)

func worker(wg *sync.WaitGroup, wid int, jobs <-chan uint64) {
	defer wg.Done()

	zap.S().Infof("Worker: %d starting", wid)
	// make db or other connect
	client, err := ethereum.Connect(config.RPC.Endpoint)
	if err != nil {
		zap.S().Error(err.Error())
	}
	blockRepo := blockRepository.NewBlockRepository(utils.DB())
	blockUsec := blockUsecase.NewBlockUsecase(blockRepo, client)

	for job := range jobs {
		zap.S().Infof("Worker: %d, started job: %d", wid, job)

		if err := blockUsec.CreateByBlockNum(job); err != nil {
			zap.S().Error(err.Error())
		}
		zap.S().Infof("Worker: %d, finished job: %d", wid, job)
	}
	zap.S().Infof("Worker: %d, interrupted", wid)
}
