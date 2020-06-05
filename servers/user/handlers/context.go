package handlers

import (
	"MeetingScheduler/servers/user/model"
	"MeetingScheduler/servers/user/sessions"
)

// HandlerCtx provides access to context for HTTP handler functions
type HandlerCtx struct {
	SigningKey   string
	SessionStore sessions.Store
	UserStore    model.Store
}
