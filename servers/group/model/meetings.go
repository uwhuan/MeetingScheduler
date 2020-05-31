package model

import "time"

// Meeting struct contains information about a meeting
type Meeting struct {
	MeetingID  int       `json:"meetingID"`
	Name       string    `json:"name"`
	Creator    int       `json:"creator"`
	Schedule   Schedule  `json:"-"`
	CreateDate time.Time `json:"createDate"`
	Confirmed  bool      `json:"confirmed"` // defalt false
}
