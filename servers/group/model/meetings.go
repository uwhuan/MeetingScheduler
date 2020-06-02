package model

// Meeting struct contains information about a meeting
type Meeting struct {
	MeetingID   int64  `json:"meetingID"`
	GroupID     int64  `json:"groupID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatorID   int64  `json:"creatorID"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	CreateDate  string `json:"createDate"`
	Confirmed   int    `json:"confirmed"` // defalt false
}
