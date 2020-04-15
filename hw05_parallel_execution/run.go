package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error { //nolint:gocritic
	wg := &sync.WaitGroup{}
	tasksCh := make(chan Task)

	var runErrorssCount int32

	for i := 0; i < N; i++ {
		go startWorker(wg, tasksCh, &runErrorssCount, M)
	}

	for _, task := range tasks {
		wg.Add(1) //nolint:gomnd
		tasksCh <- task
	}

	close(tasksCh)
	wg.Wait()

	if atomic.LoadInt32(&runErrorssCount) > int32(M) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func startWorker(wg *sync.WaitGroup, tasks <-chan Task, runErrorssCount *int32, maxErrors int) {
	for task := range tasks {
		if atomic.LoadInt32(runErrorssCount) <= int32(maxErrors) {
			err := task()
			if err != nil {
				atomic.AddInt32(runErrorssCount, 1)
			}
		}

		wg.Done()
	}
}
