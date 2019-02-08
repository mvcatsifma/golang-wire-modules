package db

import "fmt"

type DbClient struct {
}

func (this *DbClient) Query(){
	fmt.Println("Querying...")
}

func NewClient() *DbClient {
	return &DbClient{}
}
