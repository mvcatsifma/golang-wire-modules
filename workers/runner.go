package workers

import (
	"sync"
	"time"
)

func Go(work func(chan bool, chan error), cleanup func() error) *Task {
	t := &Task{
		stopChan: make(chan bool, 1),
		errChan:  make(chan error, 1),
	}
	go func() {
		defer func() {
			err := cleanup()
			if err != nil {
				select {
				case t.errChan <- err:
				case <-time.Tick(1 * time.Second):
				}
			}
			t.wg.Done()
		}()

		t.wg.Add(1)
		work(t.stopChan, t.errChan)
	}()
	return t
}

type Task struct {
	wg       sync.WaitGroup
	stopChan chan bool
	errChan  chan error
	err      error
}

func (t *Task) Stop() {
	t.stopChan <- true
	t.wg.Wait()
	close(t.errChan)
	close(t.stopChan)
}

func (t *Task) Errors() <-chan error {
	return t.errChan
}
