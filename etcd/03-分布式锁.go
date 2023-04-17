package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:12379", "127.0.0.1:22379", "127.0.0.1:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// 创建两个单独会话演示锁竞争
	s1, err := concurrency.NewSession(client)
	if err != nil {
		log.Fatalln(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/my-lock/")

	s2, err := concurrency.NewSession(client)
	if err != nil {
		log.Fatalln(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/my-lock/")

	// 会话s1获取锁
	if err = m1.Lock(context.Background()); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// 阻塞 直到s1释放锁
		if err = m2.Lock(context.Background()); err != nil {
			log.Fatalln(err)
		}
	}()
	if err = m1.Unlock(context.Background()); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("released lock for s1")

	<-m2Locked // 此时确保s2上锁
	fmt.Println("acquired lock for s2")
	if err = m2.Unlock(context.Background()); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("released lock for s2")
}
