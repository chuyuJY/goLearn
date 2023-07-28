package query

import (
	"trlab-backend-go/entity/chatgpt"
	"trlab-backend-go/logger"
	"trlab-backend-go/util"

	"go.uber.org/zap"
)

func InsertChatSession(chatSession *chatgpt.ChatgptSession) {
	err := Db().Model(&chatgpt.ChatgptSession{}).Create(chatSession).Error

	if err != nil {
		logger.ZapLogger.Error(util.LogChatgptTable()+"insert chatgpt session failed", zap.Error(err))
		util.ChatgptNotFound()
	}
	return
}

func GetChatSession(sessionId string) (sessions chatgpt.ChatgptSessions) {
	err := Db().Model(&chatgpt.ChatgptSession{}).Where(&chatgpt.ChatgptSession{SessionId: sessionId}).
		Order("seq DESC").Limit(20).Find(&sessions).Error
	if err != nil {
		logger.ZapLogger.Error(util.LogChatgptTable()+"get chatgpt session failed", zap.String("sessionId", sessionId), zap.Error(err))
		util.ChatgptNotFound()
	}

	return sessions
}
