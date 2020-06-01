package handler

import "MeetingScheduler/servers/group/dao"

//Context stores the userID, sessionID and SQL database
type Context struct {
	uid   int64
	sid   int64
	store *dao.Store
}
