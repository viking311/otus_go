package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

var errCount int

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n == 0 {
		return nil
	}

	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	errCount = m
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	taskCh := make(chan Task, len(tasks))

	for range n {
		wg.Add(1)
		go executor(taskCh, &wg, &mx)
	}

	for _, ts := range tasks {
		taskCh <- ts
	}
	close(taskCh)

	wg.Wait()

	if errCount <= 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func executor(ch chan Task, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()

	for ts := range ch {
		mx.Lock()
		ec := errCount
		mx.Unlock()

		if ec > 0 {
			err := ts()
			if err != nil {
				mx.Lock()
				errCount--
				mx.Unlock()
			}
		} else {
			return
		}
	}
}
