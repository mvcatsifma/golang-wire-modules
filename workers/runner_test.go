package workers

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	workFn := func() error {
		fmt.Println("Doing work...")
		time.Sleep(1 * time.Second)

		if time.Now().Unix()%2 == 0 {
			return errors.New("error")
		}
		return nil
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

	time.Sleep(20 * time.Second)

	task.Stop()
}
