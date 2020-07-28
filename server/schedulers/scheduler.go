package schedulers

import (
	commonRepo "assignment/entities/repos/common"
	cancel_odd "assignment/server/schedulers/cancel-odd"
)

func InitSchedulers(cancelOddConfig cancel_odd.Config, commonRepository commonRepo.IRepository) {
	cancelOdd := cancel_odd.New(&cancelOddConfig, commonRepository)
	go cancelOdd.Exec()
}
