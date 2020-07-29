package cancel_odd

import (
	commonRepo "assignment/entities/repos/common"
	"assignment/server/schedulers/scheduler"
	log "github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Period         int `default:"1"`
	NumberToCancel int `default:"10"`
}

type CancelOddScheduler struct {
	config *Config
	repo   commonRepo.IRepository
}

func (scheduler CancelOddScheduler) Exec() error {
	tick := time.Tick(time.Duration(scheduler.config.Period) * time.Minute)
	logger := log.WithField("scheduler", "cancelOdd")
	logger.Info("initializing")
	for range tick {
		if err := scheduler.repo.CancelOddTransactions(scheduler.config.NumberToCancel); err != nil {
			logger.WithError(err).Error("error")
		} else {
			logger.Info("executed")
		}
	}
	return nil
}

func New(config *Config, repository commonRepo.IRepository) scheduler.IScheduler {
	return &CancelOddScheduler{
		config: config,
		repo:   repository,
	}
}
