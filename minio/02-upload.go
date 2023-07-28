package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/tags"
)

func main() {
	//UploadFile(bucketName, "content_embedding.json", "./minio/content_embedding.json")
	//GetObject()
	//StatObject()
	//GetObjectTagging()
	//ListObjects()
	//IfModified()
	//RemoveObjects()
}

func UploadFile(bucketName, objectName, filePath string) {
	uploadInfo, err := minioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: "application/json",
		UserTags:    map[string]string{"LastModified": time.Now().Format(time.RFC3339)},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded object: ", uploadInfo)
}

func IfModified() {
	objectInfo, err := minioClient.StatObject(context.Background(), bucketName, "2023-07-11 18:36:45.19565 +0800 CST m=+2.777323876", minio.StatObjectOptions{})
	if err != nil {
		fmt.Println(minio.ToErrorResponse(err).Error())
		return
	}
	fmt.Println(objectInfo)
}

func ListObjects() {
	multiPartObjectCh := minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{
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

func SetUpdateTime() {
	ts, _ := tags.MapToBucketTags(map[string]string{"UpdateTime": time.Now().Format(time.RFC3339)})
	if err := minioClient.SetBucketTagging(context.Background(), bucketName, ts); err != nil {

	}

}

func GetUpdateTime() string {
	ts, err := minioClient.GetBucketTagging(context.Background(), bucketName)
	if err != nil {

	}
	return ts.ToMap()["UpdateTime"]
}

func GetObject() {
	object, err := minioClient.GetObject(context.Background(), bucketName, "Who-is-Rhizome.json", minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer object.Close()

	localFile, err := os.Create("./minio/get.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer localFile.Close()

	if _, err = io.Copy(localFile, object); err != nil {
		fmt.Println(err)
		return
	}
}

func StatObject() {
	objInfo, err := minioClient.StatObject(context.Background(), bucketName, "content_embedding.json", minio.StatObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(objInfo.UserTags["LastModified"])
}

func GetObjectTagging() {
	tags, err := minioClient.GetObjectTagging(context.Background(), bucketName, objectName, minio.GetObjectTaggingOptions{})
	if err != nil {
		errResponse := minio.ToErrorResponse(err)
		switch errResponse.Code {
		case "NoSuchKey":
			fmt.Println(errResponse.Error())
		default:
			fmt.Println(errResponse.Error())

		}
		return
	}
	fmt.Printf("Fetched Tags: %s", tags.ToMap())
}

func RemoveObjects() {
	objectsCh := make(chan minio.ObjectInfo)

	// Send object names that are needed to be removed to objectsCh
	go func() {
		defer close(objectsCh)
		// List all objects from a bucket-name with a matching prefix.
		for object := range minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{}) {
			if object.Err != nil {
				log.Fatalln(object.Err)
			}
			objectsCh <- object
		}
	}()

	opts := minio.RemoveObjectsOptions{
		GovernanceBypass: true,
	}

	for rErr := range minioClient.RemoveObjects(context.Background(), bucketName, objectsCh, opts) {
		fmt.Println("Error detected during deletion: ", rErr)
	}
}
