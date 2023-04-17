package main

import (
	"context"
	"fmt"
	"log"

	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func testLease(client *clientv3.Client) {
	// 1. 直接创建一个10s的租约
	//grantResp, err := client.Grant(context.Background(), 10)
	// 2. 获取Lease对象
	lease := clientv3.NewLease(client)
	// 创建一个10s的租约
	grantResp, err := lease.Grant(context.Background(), 10)
	if err != nil {
		log.Fatalln(err)
		return
	}
	// 10s后，/woqu/ 这个key就会被移除
	_, err = client.Put(context.Background(), "/woqu/", "quququ", clientv3.WithLease(grantResp.ID))
	if err != nil {
		log.Fatalln(err)
		return
	}
	// KeepAlive 续约
	// 有协程来帮自动续租，每秒一次。
	keepAliveChan, err := client.KeepAlive(context.Background(), grantResp.ID)
	if err != nil {
		log.Println(err)
	}
	for {
		ka := <-keepAliveChan
		log.Printf("ttl:%v\n", ka.TTL)
	}
}

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:12379", "127.0.0.1:22379", "127.0.0.1:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("Connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("Connect to etcd success...")
	defer client.Close()
	// lease test
	testLease(client)
}
