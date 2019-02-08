package main

import (
	"github.com/mvcatsifma/golang-wire-modules/products"
	"os"
	"os/signal"
)

func main() {
	module := products.BuildModule()
	module.Start()

	// create a channel to receive incoming OS interrupts (such as Ctrl-C):
	osInterruptChannel := make(chan os.Signal, 1)
	signal.Notify(osInterruptChannel, os.Interrupt)

	// block execution until an OS signal (such as Ctrl-C) is received:
	<-osInterruptChannel

	module.Terminate()
}

type IModule interface {
	Start()
	Terminate()
}
