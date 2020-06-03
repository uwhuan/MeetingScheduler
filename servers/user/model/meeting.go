package model

import (
	"time"
)

//Meeting represents a meeting in the database
type Meeting struct {
	MeetingID   int64     `json:"mid"`
	Name        string    `json:"name"`
	CreatorID   int64     `json:"cid"`
	GroupID     int64     `json:"gid"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	CreateDate  time.Time `json:"createDate"`
	Description string    `json:"description"`
	Confirmed   bool      `json:"confirmed"`
}

//MeetingParticipant represents a meeting participant in the database
type MeetingParticipant struct {
	MeetingID int64 `json:"mid"`
	UID       int64 `json:"uid"`
}

//MeetingParticipant represents a meeting participant in the database
type MeetingDetails struct {
	Meeting      Meeting `json:"meeting"`
	Participants []int64 `json:"participants"`
}
