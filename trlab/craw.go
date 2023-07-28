package util

import (
	"encoding/json"
	"trlab-backend-go/logger"
	"trlab-backend-go/payload"

	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

func GetStrapiArticles(url string) (articles payload.StrapiArticles) {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		logger.ZapLogger.Info("GetStrapiArticles", zap.String("url", r.URL.String()))
	})
	c.OnResponse(func(response *colly.Response) {
		if err := json.Unmarshal(response.Body, &articles); err != nil {
			logger.ZapLogger.Error("get strapi articles failed", zap.Error(err))
			return
		}
	})
	_ = c.Visit(url)
	return
}
