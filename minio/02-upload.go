package main

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func main() {
	UploadFile("hello", "content_embedding.json", "./minio/content_embedding.json")
}

func UploadFile(bucketName, objectName, filePath string) {
	uploadInfo, err := minioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: "application/json",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded object: ", uploadInfo)
}

func ListObjects() {
	multiPartObjectCh := minioClient.ListObjects(context.Background(), "hello", minio.ListObjectsOptions{
		Prefix:    "",
		Recursive: true,
	})
	for multiPartObject := range multiPartObjectCh {
		if multiPartObject.Err != nil {
			fmt.Println(multiPartObject.Err)
			return
		}
		fmt.Println(multiPartObject.Key)
	}
}
