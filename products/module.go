package products

import (
	"fmt"
	"sync"
	"time"
)

type Module struct {
	Api         *Api
	wg          sync.WaitGroup
	terminators []chan<- bool
}

func NewModule(api *Api) *Module {
	return &Module{Api: api}
}

func (this *Module) Start() {
	fmt.Println("Starting...")
	this.wg = sync.WaitGroup{}
	this.terminators = append(this.terminators, worker1(this.wg))
	this.terminators = append(this.terminators, worker2(this.wg))
}

func (this *Module) Terminate() {
	fmt.Println("Terminating...")
	for _, terminate := range this.terminators {
		terminate <- true
	}
	this.wg.Wait()
}

func worker1(wg sync.WaitGroup) (terminate chan bool) {
	terminate = make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-terminate:
				fmt.Println("worker1: terminating")
				time.Sleep(3) // do expensive cleanup work
				fmt.Println("worker1: terminated")
				return
			default:
				time.Sleep(2 * time.Second)
				fmt.Println("worker1: working")
			}
		}
	}()
	return
}

func worker2(wg sync.WaitGroup) (terminate chan bool) {
	terminate = make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-terminate:
				fmt.Println("worker2: terminating")
				time.Sleep(2) // do expensive cleanup work
				fmt.Println("worker2: terminated")
				return
			default:
				time.Sleep(5 * time.Second)
				fmt.Println("worker2: working")
			}
		}
	}()
	return
}
