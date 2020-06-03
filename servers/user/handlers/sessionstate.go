package handlers

import (
	"time"

	"github.com/my/repo/ljchen17/final-project/MeetingScheduler/servers/user/model"
)

// SessionState represents a session for an authenticated user
type SessionState struct {
	BeginTime time.Time `json:"beginTime"`
	User      *model.User
}
