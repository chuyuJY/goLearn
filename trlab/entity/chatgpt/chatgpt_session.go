package chatgpt

import (
	"trlab-backend-go/util/db"
)

type ChatgptSessions []ChatgptSession

type ChatgptSession struct {
	SessionId string
	Seq       uint
	Did       string `gorm:"default:null"`
	Role      string
	Context   string
	Time      string
}

func (ChatgptSession) TableName() string {
	return db.ChatgptSessionTable
}
