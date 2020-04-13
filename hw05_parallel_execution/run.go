package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error {
	wg := &sync.WaitGroup{}
	tasksCh := make(chan Task)

	for i := 0; i < N; i++ {
		go startWorker(wg, tasksCh)
	}

	for _, task := range tasks {
		wg.Add(1)
		tasksCh <- task
	}

	wg.Wait()
	close(tasksCh)
	return nil
}

func startWorker(wg *sync.WaitGroup, tasks <-chan Task) {
	for task := range tasks {
		task()
		wg.Done()
	}
}
