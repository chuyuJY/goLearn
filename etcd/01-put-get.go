package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// etcd client put/get demo
// use etcd/clientv3

func putEtcd(client *clientv3.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 1. 直接通过client调用Put
	//putResp, err := client.Put(ctx, "/hello", "world")
	// 2. 通过KV调用Put
	kv := clientv3.NewKV(client)
	putResp, err := kv.Put(ctx, "/hello/first", "world")
	if err != nil {
		fmt.Printf("Put to etcd failed, err:%v\n", err)
		return
	}
	fmt.Printf("Put to etcd success, the putResp header is:%v\n", putResp.Header)
	// 再写入两个 用于后续演示
	_, err = kv.Put(ctx, "/hello/second", "zjy")
	_, err = kv.Put(ctx, "/hello/third", "tiam")
	_, err = kv.Put(ctx, "/helloxxxxx", "干扰")
	// 更新/helloxxxxx的值，并返回之前的值
	op := clientv3.OpPut("/helloxxxxx", "new干扰", clientv3.WithPrevKV())
	resp, _ := kv.Do(ctx, op)
	fmt.Printf("op success:%v\n", resp.Get().Kvs)
}

func getEtcd(client *clientv3.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 1. 直接通过client调用Get
	//resp, err := client.Get(ctx, "/hello")
	// 2. 通过KV对象调用Get
	kv := clientv3.NewKV(client)
	resp, err := kv.Get(ctx, "/hello/", clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("Get from etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("Get from etcd success, the data is:")
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
	// 开启事务
	txn := kv.Txn(ctx)
	// 如果/hello/first的值为world, 则获取/hello/second的值, 否则获取/hello/third的值。
	txnResp, err := txn.If(clientv3.Compare(clientv3.Value("/hello/first"), "=", "world")).
		Then(clientv3.OpGet("/hello/second")).
		Else(clientv3.OpGet("/hello/third")).
		Commit()
	// 输出事务执行结果
	if txnResp.Succeeded { // If == true
		fmt.Println("~~~", txnResp.Responses[0].GetResponseRange().Kvs)
	} else {
		fmt.Println("!!!", txnResp.Responses[0].GetResponseRange().Kvs)
	}
}

func watchEtcd(client *clientv3.Client) {
	// watch key:/helloxxxxx change
	watchChan := client.Watch(context.Background(), "/helloxxxxx") // <-chan WatchResponse
	// 1. 读取变化
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", event.Type, event.Kv.Key, event.Kv.Value)
		}
	}
	// 2. 读取变化
	//select {
	//case <-watchChan:
	//	fmt.Println(client.Get(context.Background(), "/helloxxxxx"))
	//}
}

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:12379", "127.0.0.1:22379", "127.0.0.1:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error
		fmt.Printf("Connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("Connect to etcd success...")
	defer client.Close()
	// put
	putEtcd(client)
	// get
	getEtcd(client)
}
