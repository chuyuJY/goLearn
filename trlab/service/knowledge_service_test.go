package service

import (
	"testing"
	"time"
	"trlab-backend-go/entity/chatgpt"
	"trlab-backend-go/payload"
	"trlab-backend-go/query"
	"trlab-backend-go/util"
	"trlab-backend-go/util/db"

	"github.com/stretchr/testify/require"
)

var milvusTestCollectionName = "trlab_test"

func TestUpdateMilvus(t *testing.T) {
	t.SkipNow()
	util.SetupTestMilvus(t)
	db.Milvus().LoadCollection(milvusTestCollectionName)
	type args struct {
		collectionName string
		milvusData     *chatgpt.MilvusData
	}
	tests := []struct {
		name string
		args args
	}{
		{"article1", args{milvusTestCollectionName, PrepareMilvusData(&query.TestChatgptArticle1)}},
		{"article2", args{milvusTestCollectionName, PrepareMilvusData(&query.TestChatgptArticle2)}},
		{"article3", args{milvusTestCollectionName, PrepareMilvusData(&query.TestChatgptArticle3)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.True(t, UpdateMilvus(tt.args.collectionName, tt.args.milvusData))
			require.True(t, DeleteEntityBySlugs(tt.args.collectionName, tt.args.milvusData.Slug))
		})
	}
}

func TestUpdateKnowledgeLibrary(t *testing.T) {
	t.SkipNow()
	util.SetupTestMilvus(t)
	db.Milvus().LoadCollection(milvusTestCollectionName)
	type args struct {
		articles []payload.Article
	}
	tests := []struct {
		name string
		args args
	}{
		{"articles", args{articles: query.TestChatgptArticles}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uploadRes1 := UpdateKnowledgeLibraryWithArticles(milvusTestCollectionName, tt.args.articles)
			require.Equal(t, uploadRes1.success, len(tt.args.articles))
			for i := 0; i < len(tt.args.articles); i++ {
				tt.args.articles[i].UpdatedAt = time.Now()
			}
			uploadRes2 := UpdateKnowledgeLibraryWithArticles(milvusTestCollectionName, tt.args.articles)
			require.Equal(t, uploadRes2.success, len(tt.args.articles))
			require.Equal(t, uploadRes2.failure, 0)
		})
	}
	for _, tt := range tests {
		for _, article := range tt.args.articles {
			require.True(t, DeleteEntityBySlugs(milvusTestCollectionName, article.Slug))
		}
	}
}
