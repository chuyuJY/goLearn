package service

import (
	"fmt"
	"regexp"
	"strings"
	"trlab-backend-go/config"
	"trlab-backend-go/entity/chatgpt"
	"trlab-backend-go/logger"
	"trlab-backend-go/payload"
	"trlab-backend-go/util"
	"trlab-backend-go/util/db"

	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"go.uber.org/zap"
)

var (
	StrapiEditorialBaseUrl = "https://strapi.trlab.com/api/editorials?populate=*&sort[0]=updatedAt:desc&pagination[start]=0&pagination[limit]=100"
)

func UpdateKnowledgeLibraryWithArticles(collectionName string, articles []payload.Article) struct {
	success int
	failure int
} {
	logger.ZapLogger.Info("update knowledge", zap.Int("total", len(articles)))
	var count int
	for _, article := range articles {
		milvusData := PrepareMilvusData(&article)
		if UpdateMilvus(collectionName, milvusData) {
			count++
		}
	}
	logger.ZapLogger.Info("update knowledge finished", zap.Int("failure", len(articles)-count), zap.Int("success", count))
	return struct {
		success int
		failure int
	}{success: count, failure: len(articles) - count}
}

func CompareIfModified(collectionName string, latestArts []payload.Article) (updateArts []payload.Article) {
	for _, article := range latestArts {
		lastModified, ok := db.Milvus().QueryLastModified(collectionName, fmt.Sprintf("slug in [\"%v\"]", article.Slug), "updatedAt")
		if !ok || lastModified.Before(article.UpdatedAt) {
			updateArts = append(updateArts, article)
		}
	}
	return
}

func PrepareMilvusData(article *payload.Article) *chatgpt.MilvusData {
	PreProcessArticle(article)
	textBot := DefaultChatBot()
	defer textBot.Close()
	milvusData := &chatgpt.MilvusData{Article: article}
	chunks := util.SafeGetChunks(textBot.Client, util.EmbeddingModel, milvusData.Content, false, util.EmbeddingLimitToken)
	for _, chunk := range chunks {
		milvusData.Chunks = append(milvusData.Chunks, *chunk)
	}

	return milvusData
}

func UpdateMilvus(collectionName string, milvusData *chatgpt.MilvusData) bool {
	logger.ZapLogger.Info("update milvus", zap.String("title", milvusData.Title))

	if DeleteEntityBySlugs(collectionName, milvusData.Slug) && db.Milvus().BulkInsert(collectionName, BuildInsertColumns(milvusData)...) {
		logger.ZapLogger.Info("update milvus success")
		return true
	} else {
		logger.ZapLogger.Info("update milvus failed")
		return false
	}
}

func DeleteEntityBySlugs(collectionName string, slugs ...string) bool {
	ids, ok := db.Milvus().QueryIds(collectionName, fmt.Sprintf("slug in [\"%v\"]", strings.Join(slugs, "\",\"")), db.MilvusChatgptPkId)
	if !ok {
		return false
	}

	if len(ids) > 0 && !db.Milvus().BulkDeleteByIds(collectionName, db.MilvusChatgptPkId, ids) {
		return false
	}
	return true
}

func BuildInsertColumns(milvusData *chatgpt.MilvusData) []entity.Column {
	rows := len(milvusData.Chunks)

	slugs := make([]string, rows)
	titles := make([]string, rows)
	authors := make([]string, rows)
	urls := make([]string, rows)
	updatedAts := make([]string, rows)
	embeddings := make([][]float32, rows)
	contents := make([]string, rows)

	for i := 0; i < rows; i++ {
		slugs[i] = milvusData.Slug
		titles[i] = milvusData.Title
		authors[i] = milvusData.Author
		urls[i] = milvusData.Url
		updatedAts[i] = milvusData.UpdatedAt.Format(util.RFC3339Mill)
		embeddings[i] = milvusData.Chunks[i].Embedding
		contents[i] = milvusData.Chunks[i].Content
	}

	slugCol := entity.NewColumnVarChar("slug", slugs)
	titleCol := entity.NewColumnVarChar("title", titles)
	authorCol := entity.NewColumnVarChar("author", authors)
	urlCol := entity.NewColumnVarChar("url", urls)
	updatedAtCol := entity.NewColumnVarChar("updatedAt", updatedAts)
	embeddingCol := entity.NewColumnFloatVector("embedding", util.EmbeddingDimension, embeddings)
	contentCol := entity.NewColumnVarChar("content", contents)

	return []entity.Column{slugCol, titleCol, authorCol, urlCol, updatedAtCol, embeddingCol, contentCol}
}

func getEditorialBaseUrl() string {
	if config.GetConfig().Stage() == "dev" {
		return "https://dev-main-site.trlab.fun/editorial/"
	} else if config.GetConfig().Stage() == "testnet" {
		return "https://testnet-main-site.trlab.fun/editorial/"
	} else if config.GetConfig().Stage() == "pre" {
		return "https://pre-main-site.trlab.fun/editorial/"
	} else {
		return "https://trlab.com/editorial/"
	}
}

func FetchStrapiArticles(lastModified string) []payload.Article {
	return GetAllStrapiArticles(fmt.Sprintf(StrapiEditorialBaseUrl+"&filters[updatedAt][$gte]=%s", lastModified))
}

func PreProcessArticle(article *payload.Article) {
	re := regexp.MustCompile("<[^>]*>")
	article.Content = re.ReplaceAllString(article.Content, "")
	article.Url = getEditorialBaseUrl() + article.Slug
}

func GetAllStrapiArticles(url string) []payload.Article {
	strapiArticles := util.GetStrapiArticles(url)
	articles := make([]payload.Article, len(strapiArticles.Data))
	for i, strapiArticle := range strapiArticles.Data {
		articles[i] = strapiArticle.Attributes
	}
	return articles
}
