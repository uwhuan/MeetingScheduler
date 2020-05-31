package model

// Meeting struct contains information about a meeting
type Meeting struct {
	MeetingID   int64  `json:"meetingID"`
	GroupID     int64  `json:"groupID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Creator     int64  `json:"creator"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	CreateDate  string `json:"createDate"`
	Confirmed   bool   `json:"confirmed"` // defalt false
}
