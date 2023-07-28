package api

import (
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"trlab-backend-go/api/middleware"
	"trlab-backend-go/logger"
	"trlab-backend-go/payload"
	"trlab-backend-go/service"
	"trlab-backend-go/util"
	"trlab-backend-go/util/db"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

// Chat Post api/v1/chatgpt/chat
func Chat(router *gin.RouterGroup) {
	router.POST("/chat",
		middleware.OptPrivyAuthMiddleware(),
		func(c *gin.Context) {
			var req payload.ChatGPTChatReq
			if err := c.ShouldBindJSON(&req); err != nil {
				util.BadRequestError(err.Error())
			}
			logger.ZapLogger.Info("user chat request", zap.Any("request", req))
			question := req.Messages[len(req.Messages)-1].Content
			chatBot := service.DefaultChatBot()
			defer chatBot.Close()

			chatBot.SetSession(c.MustGet("Session-Id").(string))
			chatBot.Session.Did = c.GetString(middleware.CDid)

			resp, err := chatBot.Query(question)
			if err != nil {
				logger.ZapLogger.Error("ChatBot query failed", zap.Error(err))
				return
			}
			go chatBot.Session.AddMessage(openai.ChatMessageRoleUser, question)

			c.Header("Session-Id", chatBot.Session.SessionId)
			if !chatBot.ChatConfig.Stream {
				chatBot.Session.AddMessage(openai.ChatMessageRoleAssistant, resp.Choices[0].Message.Content)
				c.JSON(http.StatusOK, resp.Choices[0].Message.Content)
				return
			}

			var builder strings.Builder
			c.Stream(func(w io.Writer) bool {
				for {
					response, err := chatBot.Receive()
					if errors.Is(err, io.EOF) {
						chatBot.Session.AddMessage(openai.ChatMessageRoleAssistant, builder.String())
						logger.ZapLogger.Info("Chatgpt stream", zap.String("answer", builder.String()))
						logger.ZapLogger.Info("Chatgpt Stream finished")
						return false
					}

					if err != nil {
						logger.ZapLogger.Error("Chatgpt stream", zap.Error(err))
						return false
					}
					if len(response.Choices) == 0 {
						logger.ZapLogger.Error("Chatgpt stream", zap.Error(err))
						return false
					}
					c.SSEvent("", response)
					content := response.Choices[0].Delta.Content
					builder.WriteString(content)
					return true
				}
			})
		})
}

// GetPrevContext Get api/v1/chatgpt/getcontext
func GetPrevContext(router *gin.RouterGroup) {
	router.GET("/getcontext",
		func(c *gin.Context) {
			session, exist := service.SessionManager.GetSession(c.GetHeader("Session-Id"))
			if !exist {
				logger.ZapLogger.Info("Session queried does not exist", zap.String("SessionId", c.GetHeader("Session-Id")))
				c.AbortWithStatusJSON(http.StatusNoContent, nil)
				return
			}
			c.AbortWithStatusJSON(http.StatusOK, session.GetMessages())
		})
}

// UpdateKnowledge Post api/v1/chatgpt/update-knowledge
func UpdateKnowledge(router *gin.RouterGroup) {
	router.POST("/update-knowledge",
		middleware.StomAuthMiddleware(),
		func(c *gin.Context) {
			var startTime string
			timestamp := c.Query("startTime")
			if timestamp == "" {
				startTime = time.Time{}.Format(util.RFC3339Mill)
			} else {
				ts, err := strconv.ParseInt(c.Query("startTime"), 10, 64)
				if err != nil {
					logger.ZapLogger.Error("update knowledge failed", zap.Error(err))
					c.Abort()
				}
				startTime = time.UnixMilli(ts).UTC().Format(util.RFC3339Mill)
			}
			articles := service.FetchStrapiArticles(startTime)
			uploadRes := service.UpdateKnowledgeLibraryWithArticles(db.MilvusChatgptCollection, articles)
			c.JSON(http.StatusOK, uploadRes)
		})
}

// UploadArticles Post api/v1/chatgpt/upload-articles
func UploadArticles(router *gin.RouterGroup) {
	router.POST("/upload-articles",
		middleware.StomAuthMiddleware(),
		func(c *gin.Context) {
			var req []payload.Article
			if err := c.ShouldBindJSON(&req); err != nil {
				util.BadRequestError(err.Error())
			}
			uploadRes := service.UpdateKnowledgeLibraryWithArticles(db.MilvusChatgptCollection, req)
			c.JSON(http.StatusOK, uploadRes)
		})
}
