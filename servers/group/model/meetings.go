package model

import "time"

// Meeting struct contains information about a meeting
type Meeting struct {
	MeetingID  int64     `json:"meetingID"`
	Name       string    `json:"name"`
	Creator    int64     `json:"creator"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	CreateDate time.Time `json:"createDate"`
	Confirmed  bool      `json:"confirmed"` // defalt false
}
