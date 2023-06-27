package main

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	endpoint        = "localhost:9000"
	accessKeyID     = "XlHwS3DRQXqBXjiCzO91"
	secretAccessKey = "KBY835ks1BOFf6gMM80CQRdFlVoJWLYH6KN2Ho8l"
	useSSL          = false
)

var minioClient *minio.Client

func init() {
	var err error
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	//ok, _ := minioClient.BucketExists(context.Background(), "hello")
	//log.Printf("%#v\n", ok)
}
