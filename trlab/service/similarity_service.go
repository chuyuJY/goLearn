package service

import (
	"fmt"
	"trlab-backend-go/util/db"
)

func GetSimilarTextFromMilvus(collectionName string, queryEmbedding []float32, topN int) ([]string, []float32) {
	var texts []string
	searchResult := db.Milvus().Search(collectionName, queryEmbedding, topN, []string{"title", "author", "url", "content"})
	scores := searchResult[0].Scores
	for i := 0; i < searchResult[0].ResultCount; i++ {
		fields := searchResult[0].Fields
		title, _ := fields.GetColumn("title").GetAsString(i)
		author, _ := fields.GetColumn("author").GetAsString(i)
		url, _ := fields.GetColumn("url").GetAsString(i)
		content, _ := fields.GetColumn("content").GetAsString(i)

		texts = append(texts, fmt.Sprintf("Title: %v Author: %v URL: `%v` Content: %v", title, author, url, content))
	}
	return texts, scores
}
