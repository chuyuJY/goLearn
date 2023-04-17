package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func sendMessage() {
	// 1. producer配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel中返回
	// 2. 消息构造
	msg := &sarama.ProducerMessage{
		Topic: "web_log",
		Value: sarama.StringEncoder("this is a test log"),
	}
	// 3. 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 4. 发送消息
	partition, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed, err:", err)
		return
	}
	fmt.Printf("partition:%v offset:%v\n", partition, offset)
}

func main() {
	sendMessage()
}
