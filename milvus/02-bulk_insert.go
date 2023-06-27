package main

import (
	"context"
	"fmt"
)

func main() {
	bulkInsert()
}

func bulkInsert() {
	_, err := milvusClient.BulkInsert(context.Background(), "hello", "", []string{"http://localhost:9000/hello/content_embedding.json"})
	if err != nil {
		fmt.Println("some err:", err)
	}
}
