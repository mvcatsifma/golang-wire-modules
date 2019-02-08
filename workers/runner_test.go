package workers

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestRunner2(t *testing.T) {
	workFn := func(stopChan chan bool, errChan chan error) {
		fmt.Println("Doing more work")
	}
	cleanupFn := func() error {
		fmt.Println("Cleaning up...")
		time.Sleep(2 * time.Second)
		return nil
	}
	task := Go(workFn, cleanupFn)

	go func() {
		for err := range task.Errors() {
			fmt.Println(err)
		}
	}()

	time.Sleep(5 * time.Second)

	task.Stop()
}

func TestRunner(t *testing.T) {
	workFn := func(stopChan chan bool, errChan chan error) {
		for {
			select {
			case <-stopChan:
				fmt.Println("Terminating")
				return
			default:
				fmt.Println("Doing work...")
				time.Sleep(1 * time.Second)
				if time.Now().Unix()%2 == 0 {
					err := errors.New("error")
					select {
					case errChan <- err:
					case <-time.Tick(1 * time.Second):
					}
				}
			}
		}
	}
	cleanupFn := func() error {
		fmt.Println("Cleaning up...")
		time.Sleep(2 * time.Second)
		return nil
	}
	task := Go(workFn, cleanupFn)

	go func() {
		for err := range task.Errors() {
			fmt.Println(err)
		}
	}()

	time.Sleep(5 * time.Second)

	task.Stop()
}
