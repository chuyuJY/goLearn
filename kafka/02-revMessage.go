package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func revMessage() {
	wg := sync.WaitGroup{}
	// 1. 创建消费者
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("start consumer failed, err:", err)
		return
	}
	defer consumer.Close()
	// 2. 取消息
	partitionList, err := consumer.Partitions("web_log") // 根据该Topic取所有的分区
	if err != nil {
		fmt.Println("get list of partition failed, err:", err)
		return
	}
	fmt.Println(partitionList)
	// 遍历所有分区
	for _, partition := range partitionList {
		// 针对每个分区创建对应的分区消费者
		/*
			参数说明：
				参数1 指定消费哪个 topic
				参数2 分区 这里输出了所有分区的，一般默认消费 0 号分区。kafka 中有分区的概念，类似于ES和MongoDB中的sharding，MySQL中的分表这种
				参数3 offset 从哪儿开始消费起走，正常情况下每次消费完都会将这次的offset提交到kafka，然后下次可以接着消费，
					这里 sarama.OffsetNewest 就从最新的开始消费，即该 consumer 启动之前产生的消息都无法被消费
					如果改为 sarama.OffsetOldest 则会从最旧的消息开始消费，即每次重启 consumer 都会把该 topic 下的所有消息消费一次
		*/
		pc, err := consumer.ConsumePartition("web_log", partition, sarama.OffsetOldest)
		if err != nil {
			fmt.Println("start partition consumer failed, err:", err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费消息
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%s Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
			wg.Done()
		}(pc)
	}
	wg.Wait()
}

func main() {
	revMessage()
}
