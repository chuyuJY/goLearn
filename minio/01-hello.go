package main

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	endpoint        = "localhost:9000"
	accessKeyID     = "minioadmin"
	secretAccessKey = "minioadmin"
	region          = "cn-north-1"
	bucketName      = "milvus-bucket"
	objectName      = "Who-is-Rhizome.json"
)

var minioClient *minio.Client

func init() {
	var err error
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Region: region,
	})
	if err != nil {
		log.Fatalln(err)
	}
	//ok, _ := minioClient.BucketExists(context.Background(), "hello")
	//log.Printf("%#v\n", ok)
}
