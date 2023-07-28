package db

import (
	"context"
	"sync"
	"time"
	"trlab-backend-go/exception"
	"trlab-backend-go/logger"

	milvus "github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"go.uber.org/zap"
)

const (
	MilvusChatgptCollection = "chatgpt"
	MilvusChatgptPkId       = "id"
)

var onceMilvus sync.Once

type MilvusHelper struct {
	milvusClient milvus.Client
}

var milvusHelper MilvusHelper

func Milvus() *MilvusHelper {
	return &milvusHelper
}

func SetMilvus(uri string, username string, password string) {
	onceMilvus.Do(func() {
		if c, err := milvus.NewDefaultGrpcClientWithURI(context.Background(), uri, username, password); err != nil {
			logger.ZapLogger.Error("milvus initialized failed", zap.Error(err))
		} else {
			milvusHelper.milvusClient = c
			Milvus().LoadCollection(MilvusChatgptCollection)
			logger.ZapLogger.Info("milvus initialized successful!")
		}
	})
}

func (mh *MilvusHelper) LoadCollection(collectionName string) {
	hasColl, err := mh.milvusClient.HasCollection(context.Background(), collectionName)
	if err != nil {
		logger.ZapLogger.Error("check whether collection exists failed", zap.Error(err))
		return
	}
	if !hasColl {
		logger.ZapLogger.Error("collection does not exist", zap.String("collectionName", collectionName))
		return
	}

	err = mh.milvusClient.LoadCollection(context.Background(), collectionName, false)
	if err != nil {
		logger.ZapLogger.Error("load collection failed", zap.Error(err))
		return
	}
}

func (mh *MilvusHelper) Search(collectionName string, queryEmbedding []float32, topN int, outputs []string) []milvus.SearchResult {
	sp, _ := entity.NewIndexFlatSearchParam()
	searchResult, err := mh.milvusClient.Search(
		context.Background(),
		collectionName,
		[]string{},
		"",
		outputs,
		[]entity.Vector{entity.FloatVector(queryEmbedding)},
		"embedding",
		entity.IP,
		topN,
		sp,
	)
	if err != nil {
		panic(&exception.TRLabError{MessageID: exception.InternalError,
			TemplateData: map[string]interface{}{"Reason": "Milvus search failed"}})
	}
	return searchResult
}

func (mh *MilvusHelper) BulkInsert(collectionName string, columns ...entity.Column) bool {
	if _, err := mh.milvusClient.Insert(context.Background(), collectionName, "", columns...); err != nil {
		logger.ZapLogger.Error("milvus bulk insert failed", zap.Error(err))
		return false
	}
	return true
}

func (mh *MilvusHelper) QueryLastModified(collectionName, expr, columnName string) (time.Time, bool) {
	queryResult, err := mh.milvusClient.Query(context.Background(), collectionName, []string{}, expr, []string{columnName})
	if err != nil {
		logger.ZapLogger.Error("milvus query last modified failed", zap.Error(err))
		return time.Now(), false
	}

	updatedAt := queryResult.GetColumn(columnName).FieldData().GetScalars().GetStringData().GetData()
	if len(updatedAt) == 0 {
		return time.Time{}, true
	}
	lastModified, err := time.Parse(time.RFC3339, updatedAt[0])
	if err != nil {
		logger.ZapLogger.Error("milvus parse time failed", zap.Error(err))
		return time.Time{}, false
	}
	return lastModified, true
}

func (mh *MilvusHelper) QueryIds(collectionName, expr, columnName string) ([]int64, bool) {
	queryResult, err := mh.milvusClient.Query(context.Background(), collectionName, []string{}, expr, []string{columnName})
	if err != nil {
		logger.ZapLogger.Error("milvus bulk query failed", zap.Error(err))
		return nil, false
	}
	ids := queryResult.GetColumn(columnName).FieldData().GetScalars().GetLongData().GetData()
	return ids, true
}

func (mh *MilvusHelper) BulkDeleteByIds(collectionName, idName string, ids []int64) bool {
	idsColumn := entity.NewColumnInt64(idName, ids)
	if err := mh.milvusClient.DeleteByPks(context.Background(), collectionName, "", idsColumn); err != nil {
		logger.ZapLogger.Error("milvus bulk delete by ids failed", zap.Error(err))
		return false
	}
	return true
}

func (mh *MilvusHelper) Close() error {
	return mh.milvusClient.Close()
}
