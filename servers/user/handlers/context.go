package handlers

import (
	"github.com/my/repo/ljchen17/final-project/MeetingScheduler/servers/user/model"
	"github.com/my/repo/ljchen17/final-project/MeetingScheduler/servers/user/sessions"
)

// HandlerCtx provides access to context for HTTP handler functions
type HandlerCtx struct {
	SigningKey   string
	SessionStore sessions.Store
	UserStore    model.Store
}
