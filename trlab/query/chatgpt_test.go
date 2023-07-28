package query

import (
	"fmt"
	"testing"
	"trlab-backend-go/util"
	"trlab-backend-go/util/db"

	"github.com/stretchr/testify/require"
)

var ChatgptCleanupQueries = []string{
	"truncate " + db.ChatgptSessionTable + " cascade",
	"truncate " + db.UserTable + " cascade",
}

func TestGetChatSession(t *testing.T) {
	util.SetupTestDB(t, ChatgptCleanupQueries, nil)
	util.RunSql(ChatgptCleanupQueries)
	// create test session
	CreateTestUser(TestUser1)
	CreateTestChatgptSession(TestChatgptSession1)
	CreateTestChatgptSession(TestChatgptSession2)
	CreateTestChatgptSession(TestChatgptSession3)
	CreateTestChatgptSession(TestChatgptSession4)
	t.Run("context in chatgpt session", func(t *testing.T) {
		sessions := GetChatSession("1")
		require.True(t, len(sessions) == 2)
		for _, session := range sessions {
			require.True(t, session.Context == fmt.Sprintf("test%d", session.Seq))
			require.Equal(t, session.Did, TestUser1.Did)
		}
		sessions = GetChatSession("2")
		require.True(t, len(sessions) == 2)
		for _, session := range sessions {
			require.True(t, session.Context == fmt.Sprintf("test%d", session.Seq))
			require.True(t, session.Did == "")
		}
	})
	util.RunSql(ChatgptCleanupQueries)
}
