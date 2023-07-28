package middleware

import (
	"errors"
	"fmt"
	"strings"
	"trlab-backend-go/logger"
	"trlab-backend-go/privy"
	"trlab-backend-go/query"
	"trlab-backend-go/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetAccessToken(c *gin.Context) (accessToken string) {
	tokenCookie, err := c.Cookie(util.AuthorKey)
	if err != nil || len(tokenCookie) == 0 {
		logger.ZapLogger.Warn("token cookie is not provided", zap.Error(err))
		c.Abort()
		return
	}

	fields := strings.Fields(tokenCookie)
	if len(fields) < 2 {
		err := errors.New("invalid authorization cookie format")
		logger.ZapLogger.Warn("invalid authorization cookie format", zap.Error(err))
		c.Abort()
		return
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != strings.ToLower(util.AuthorTypeBearer) {
		err := fmt.Errorf("unsupported authorization type %s", authorizationType)
		logger.ZapLogger.Warn("get jwt token failed", zap.Error(err))
		c.Abort()
		return
	}

	accessToken = fields[1]
	return
}

const CDid = "did"
const CUid = "uid"

// AuthMiddleware creates a gin privy middleware for authorization
func PrivyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := GetAccessToken(c)
		payload, err := privy.ParseToken(accessToken)
		if err != nil {
			logger.ZapLogger.Warn("verify privy token failed", zap.Error(err), zap.String("accessToken", accessToken))
			util.UserSessionExpired(c)
			return
		}
		c.Set(CDid, payload.UserId)
		id := query.DidToId(payload.UserId)
		c.Set(CUid, id)
		c.Next()
	}
}

// OptPrivyAuthMiddleware creates a gin optional privy middleware for authorization
func OptPrivyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := GetAccessToken(c)
		if len(accessToken) == 0 {
			c.Set(CDid, "")
		} else {
			payload, err := privy.ParseToken(accessToken)
			if err != nil {
				c.Set(CDid, "")
			} else {
				c.Set(CDid, payload.UserId)
			}
		}
		c.Next()
	}
}
