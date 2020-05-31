package model

import "time"

// ScheduleID is the id of schedule
type ScheduleID int

// Schedule struct stores a slot of time for a user for a meeting.
// It is used to determine the final meeting time
type Schedule struct {
	SID       ScheduleID `json:"scheduleID"`
	StartTime time.Time  `json:"startTime"`
	EndTime   time.Time  `json:"endTime"`
}

// Votes countes the number vote for certain schedule
type Votes struct {
	SID   ScheduleID `json:"scheduleID"`
	Count int        `json:"Count"`
}
