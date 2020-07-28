package cancel_odd

import (
	commonRepo "assignment/entities/repos/common"
	"assignment/server/schedulers/scheduler"
	log "github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Period         int `default:"1"`
	NumberToCancel int `default:"0"`
}

type CancelOddScheduler struct {
	config *Config
	repo   commonRepo.IRepository
}

func (scheduler CancelOddScheduler) Exec() error {
	tick := time.Tick(time.Duration(scheduler.config.Period) * time.Minute)
	for range tick {
		log.Info("start cancel odd scheduler")
		if err := scheduler.repo.CancelOddTransactions(scheduler.config.NumberToCancel); err != nil {
			log.WithError(err).Error("cancel odd scheduler error")
		} else {
			log.Info("cancel odd scheduler success")
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
