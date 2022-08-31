package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func initDB() error {
	// 设置客户端连接
	clientOptions := options.Client().ApplyURI("mongodb://root:123456@127.0.0.1:27017")
	// 连接到db
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
		return err
	} else {
		fmt.Printf("client: %v\n", c)
	}
	err = c.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalln(err)
		return err
	} else {
		fmt.Println("连接成功!")
	}
	client = c
	return nil
}

func connectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, err
	}
	return client.Database(name), nil
}

func mongoClose() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

// 查找name字段与’张三’或’李四’匹配的文档
//func search() {
//	d := bson.D{{
//		"name",
//		bson.D{{
//			"$in",
//			bson.A{"张三", "李四"},
//		}},
//	}}
//}

//func main() {
//	err := initDB()
//	if err != nil {
//		log.Fatalln("连接失败!")
//	}
//	//
//
//}
