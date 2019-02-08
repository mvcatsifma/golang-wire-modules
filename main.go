package main

import "github.com/mvcatsifma/golang-wire-modules/products"

func main() {
	service := products.InitializeService()
	service.DoGet()
}
