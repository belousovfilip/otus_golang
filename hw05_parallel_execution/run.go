package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

var countErrors int

var (
	mx   sync.Mutex
	once sync.Once
)

var done chan interface{}

func createProducerTasks(tasks []Task) <-chan Task {
	producer := make(chan Task, len(tasks))
	go func() {
		defer close(producer)
		for _, task := range tasks {
			producer <- task
		}
	}()
	return producer
}

func consumerTasks(producer <-chan Task, maxErrors int) {
	for {
		select {
		case task, ok := <-producer:
			if !ok {
				return
			}
			mx.Lock()
			isMaxErrors := countErrors > maxErrors
			mx.Unlock()
			if isMaxErrors {
				once.Do(func() { close(done) })
				return
			}
			if err := task(); err != nil {
				mx.Lock()
				countErrors++
				mx.Unlock()
			}
		case <-done:
			return
		}
	}
}

func Run(tasks []Task, workers, maxError int) error {
	countErrors = 0
	mx = sync.Mutex{}
	once = sync.Once{}
	wg := sync.WaitGroup{}
	wg.Add(workers)
	done = make(chan interface{})
	producer := createProducerTasks(tasks)
	for ; workers > 0; workers-- {
		go func() {
			consumerTasks(producer, maxError)
			wg.Done()
		}()
	}
	wg.Wait()
	if countErrors > maxError {
		return ErrErrorsLimitExceeded
	}
	return nil
}
