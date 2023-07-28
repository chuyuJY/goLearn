package service

import (
	"sort"
	"time"
	"trlab-backend-go/entity/chatgpt"
	"trlab-backend-go/logger"
	"trlab-backend-go/query"
	"trlab-backend-go/util"

	"github.com/bwmarrin/snowflake"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
)

var SessionManager = newSessionManager()

type Session struct {
	SessionId  string
	Did        string
	SeqCounter uint
	Messages   chatgpt.ChatgptSessions
}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) GetSeq() uint {
	s.SeqCounter++
	seq := s.SeqCounter
	return seq
}

func (s *Session) AddMessage(role string, context string) {
	message := &chatgpt.ChatgptSession{
		SessionId: s.SessionId,
		Seq:       s.GetSeq(),
		Did:       s.Did,
		Role:      role,
		Context:   context,
		Time:      time.Now().Format(util.RFC3339Mill),
	}
	query.InsertChatSession(message)
	SessionManager.UpdateSession(s.SessionId, message)
}

func (s *Session) GetMessages() chatgpt.ChatgptSessions {
	return s.Messages
}

type sessionManager struct {
	IdGenerator *snowflake.Node
	Cache       *cache.Cache
}

func newSessionManager() *sessionManager {
	c := cache.New(time.Hour, 10*time.Minute)
	node, _ := snowflake.NewNode(1) // note: 部署在多节点上时，需要修改为不同 nodeId
	return &sessionManager{Cache: c, IdGenerator: node}
}

func (sm *sessionManager) SetSession(session *Session) {
	maxCacheTime := time.Hour
	sm.Cache.Set(session.SessionId, session, maxCacheTime)
}

func (sm *sessionManager) CreateSession() *Session {
	session := &Session{SessionId: sm.IdGenerator.Generate().String()}
	sm.SetSession(session)
	return session
}

func (sm *sessionManager) GetSession(sessionId string) (*Session, bool) {
	if session, exist := sm.Cache.Get(sessionId); exist {
		return session.(*Session), true
	}
	if chatgptSessions := query.GetChatSession(sessionId); len(chatgptSessions) > 0 {
		session := &Session{SessionId: sessionId}
		sort.Slice(chatgptSessions, func(i, j int) bool {
			return chatgptSessions[i].Seq < chatgptSessions[j].Seq
		})
		session.Messages = chatgptSessions
		session.SeqCounter = session.Messages[len(session.Messages)-1].Seq + 1
		sm.SetSession(session)
		return session, true
	}
	return nil, false
}

func (sm *sessionManager) UpdateSession(sessionId string, message *chatgpt.ChatgptSession) {
	session, exist := sm.GetSession(sessionId)
	if !exist {
		logger.ZapLogger.Warn("sessionManager update session", zap.String("result", "fail"))
		return
	}
	session.Messages = append(session.Messages, *message)
	sm.SetSession(session)
}

func (sm *sessionManager) RemoveSession(sessionId string) {
	sm.Cache.Delete(sessionId)
}
