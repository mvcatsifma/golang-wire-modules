package db

import "fmt"

type DbClient struct {
}

func (this *DbClient) Get(){
	fmt.Println("Getting...")
}

func NewClient() *DbClient {
	return &DbClient{}
}
