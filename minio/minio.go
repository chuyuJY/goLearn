package db

import (
	"bytes"
	"context"
	"errors"
	"sync"
	"trlab-backend-go/exception"
	"trlab-backend-go/logger"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/tags"
)

const (
	MINIO_OBJECT_NOT_EXISTS = "OBJECT_NOT_EXISTS"
)

var onceMinio sync.Once

type MinioHelper struct {
	minioClient *minio.Client
}

var minioHelper MinioHelper

func Minio() *MinioHelper {
	return &minioHelper
}

func SetMinio(endpoint, accessKeyID, secretAccessKey string) {
	var err error
	onceMinio.Do(func() {
		minioHelper.minioClient, err = minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Region: "cn-north-1",
		})
		if err != nil {
			logger.ZapLogger.Error("minio initialized failed!")
		}

	})
}

func (mh *MinioHelper) IfBucketExists(bucketName string) bool {
	if exist, err := mh.minioClient.BucketExists(context.Background(), bucketName); !exist || err != nil {
		logger.ZapLogger.Error("minio bucket not exists,  err: " + err.Error())
		panic(&exception.TRLabError{MessageID: exception.InternalError})
		return false
	}
	return true
}

func (mh *MinioHelper) GetObjectTagging(bucketName, objectName string) (*tags.Tags, error) {
	tags, err := mh.minioClient.GetObjectTagging(context.Background(), bucketName, objectName, minio.GetObjectTaggingOptions{})
	if err != nil {
		errResponse := minio.ToErrorResponse(err)
		switch errResponse.Code {
		case "NoSuchKey":
			return nil, errors.New(MINIO_OBJECT_NOT_EXISTS)
		default:
			logger.ZapLogger.Error(errResponse.Error())
			return nil, errResponse
		}
	}
	return tags, nil
}

func (mh *MinioHelper) PutObject(bucketName, objectName string, buf *bytes.Buffer, options *minio.PutObjectOptions) bool {
	// 存在同名的直接覆盖
	if _, err := mh.minioClient.PutObject(context.Background(), bucketName, objectName, buf, int64(buf.Len()), *options); err != nil {
		logger.ZapLogger.Error("minio put object failed, err: " + err.Error())
		return false
	}
	return true
}

// CompareIfModified through minio
//func CompareIfModified(latestArts Articles) (updateArts Articles) {
//	for _, article := range latestArts.Data {
//		article.Attributes.UpdatedAt = time.Now()
//		fmt.Println(article.Attributes.UpdatedAt.Format(time.RFC3339))
//		tags, err := db.Minio().GetObjectTagging(bucketName, article.Attributes.Slug+".json")
//		if err != nil {
//			if err.Error() == db.MINIO_OBJECT_NOT_EXISTS {
//				updateArts.Data = append(updateArts.Data, article)
//			}
//			continue
//		}
//
//		lastModified, err := time.Parse(time.RFC3339, tags.ToMap()["LastModified"])
//		if err != nil {
//			logger.ZapLogger.Error(err.Error())
//			return
//		}
//		if lastModified.Before(article.Attributes.UpdatedAt) {
//			updateArts.Data = append(updateArts.Data, article)
//		}
//	}
//	fmt.Println(updateArts)
//	return
//}
