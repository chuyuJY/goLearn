package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Student struct {
	Name string
	Age  int
}

// 1. 插入单条document
func insertOne(s Student, collection *mongo.Collection) (bool, error) {
	insertResult, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Insert a single document:", insertResult.InsertedID)
	return true, nil
}

// 2. 插入多条document
func insertMany(ss []interface{}, collection *mongo.Collection) (bool, error) {
	insertManyResult, err := collection.InsertMany(context.TODO(), ss)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Insert multiple documents:", insertManyResult.InsertedIDs)
	return true, nil
}

func main() {
	s1 := Student{"小红", 12}
	s2 := Student{"小兰", 10}
	s3 := Student{"小黄", 11}

	// 1. 初始化数据库
	err := initDB()
	if err != nil {
		log.Fatalln(err)
	}
	collection := client.Database("go_db").Collection("student")
	// 2. 插入单条document
	ok, err := insertOne(s1, collection)
	if !ok || err != nil {
		log.Fatalln(err)
	}
	// 3. 插入多条document
	ss := []interface{}{s2, s3}
	ok, err = insertMany(ss, collection)
	if !ok || err != nil {
		log.Fatalln(err)
	}
}
