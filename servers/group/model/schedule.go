package model

import "time"

// Schedule struct stores a slot of time for a user for a meeting.
// It is used to determine the final meeting time
type Schedule struct {
	ScheduleID int       `json:"scheduleID"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	Votes      int       `json:"votes"`
}
