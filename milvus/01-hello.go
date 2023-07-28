package main

import (
	"context"
	"log"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

var milvusClient client.Client

func init() {
	var err error
	milvusClient, err = client.NewDefaultGrpcClientWithURI(context.Background(), "127.0.0.1:19530", "", "")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("success")
	}

	if exist, err := milvusClient.HasCollection(context.Background(), "chatgpt"); !exist || err != nil {
		log.Println(err)
	} else {
		log.Println("chatgpt is exist")
	}
}
