package scheduler

type IScheduler interface {
	Exec() error
}
