package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var (
		completedTasks int
		errCount       int
		wg             sync.WaitGroup
		mu             sync.Mutex
	)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				mu.Lock()
				if completedTasks >= len(tasks) || errCount >= m {
					mu.Unlock()
					return
				}
				taskIdx := completedTasks
				completedTasks++
				mu.Unlock()
				err := tasks[taskIdx]()
				if err != nil {
					mu.Lock()
					errCount++
					mu.Unlock()
					if errCount >= m {
						return
					}
				}
			}
		}()
	}
	wg.Wait()
	if errCount >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
