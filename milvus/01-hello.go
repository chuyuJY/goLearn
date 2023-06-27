package main

import (
	"context"
	"log"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

var milvusClient client.Client

func init() {
	var err error
	milvusClient, err = client.NewDefaultGrpcClientWithURI(context.Background(), "https://in01-e11ef590049c1c3.aws-us-west-2.vectordb.zillizcloud.com:19535", "chuyu", "19990404Aa_")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("success")
	}
	defer milvusClient.Close()

	if exist, err := milvusClient.HasCollection(context.Background(), "hello"); !exist || err != nil {
		log.Println(err)
	} else {
		log.Println("hello is exist")
	}
}
