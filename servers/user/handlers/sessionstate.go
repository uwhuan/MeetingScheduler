package handlers

import (
	"MeetingScheduler/servers/user/model"
	"time"
)

// SessionState represents a session for an authenticated user
type SessionState struct {
	BeginTime time.Time `json:"beginTime"`
	User      *model.User
}
