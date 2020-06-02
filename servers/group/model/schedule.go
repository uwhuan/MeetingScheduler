package model

// Schedule struct stores a slot of time for a user for a meeting.
// It is used to determine the final meeting time
type Schedule struct {
	ScheduleID int    `json:"scheduleID"`
	MeetingID  int    `json:"meetingID"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	Votes      int    `json:"votes"`
}
