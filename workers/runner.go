package workers

import (
	"fmt"
	"sync"
	"time"
)

func Go(work func() error, cleanup func() error) *Task {
	t := &Task{
		stopChan: make(chan bool),
		errChan:  make(chan error, 1),
	}
	go func() {
		defer func() {
			t.wg.Add(1)
			err := cleanup()
			if err != nil {
				select {
				case t.errChan <- err:
				case <-time.Tick(1 * time.Second):
				}
			}
			t.wg.Done()
		}()
		for {
			select {
			case <-t.stopChan:
				fmt.Println("Terminating")
				return
			default:
				t.wg.Add(1)
				err := work()
				if err != nil {
					select {
					case t.errChan <- err:
					case <-time.Tick(1 * time.Second):
					}
				}
				t.wg.Done()
			}
		}
	}()
	return t
}

type Task struct {
	wg       sync.WaitGroup
	lock     sync.RWMutex
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
