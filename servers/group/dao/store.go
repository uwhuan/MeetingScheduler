package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	model "MeetingScheduler/servers/group/model"
)

// Store is a struct of sql dababase
type Store struct {
	Db *sql.DB
}

//ErrTargetNotFound is returned when the target can't be found
var ErrTargetNotFound = errors.New("Target not found")

//MeetingStore represents a store for Meetings
type MeetingStore interface {
	//GetMeetingByID returns the meeting with the given ID
	GetMeetingByID(id int64) (*model.Meeting, error)

	//GetAllMeetingsOfGroup returns all the meeting with the given groupID
	GetAllMeetingsOfGroup(id int64) ([]*model.Meeting, error)

	//Insert inserts the meeting into the database, and returns
	//the newly-inserted meetingID, complete with the DBMS-assigned ID
	InsertMeeting(meeting *model.Meeting) (int64, error)

	//Update applies updates to the given meeting  ID
	//and returns any errors if exist
	UpdateMeeting(id int64, update *model.Meeting) error

	//Delete deletes the meeting with the given ID
	DeleteMeeting(id int64) error

	// Confirm set the confirmed start and end time of a meeting
	// and set the confirmed flag to be true
	ConfirmMeeting(id int64, schedule *model.Schedule) error

	//GetAllParticipants get all participants of a meeting
	GetAllParticipants(meetingID int64) ([]*model.User, error)
}

//GroupStore represents a store for Groups
type GroupStore interface {
	//GetMeeingByID returns the Group with the given ID
	GetGroupByID(id int64) (*model.Group, error)

	//InsertGroup inserts the Group into the database, and returns
	//the newly-inserted GroupID, complete with the DBMS-assigned ID
	InsertGroup(group *model.Group) (int64, error)

	//UpdateGroup applies updates to the given Group  ID
	//and returns any errors if exist
	UpdateGroup(update *model.Group) error

	//DeleteGroup deletes the Group with the given ID
	DeleteGroup(id int64) error

	//GetAllMembers gets all users in the current group
	GetAllMembers(id int64) ([]*model.User, error)

	//GetAllGroups gets all groups in db
	GetAllGroups() ([]*model.Group, error)
}

//ScheduleStore represents a store for Schedules
type ScheduleStore interface {
	//GetAllSchedule returns all schedules under a meeting
	GetAllSchedule(meetingID int64) ([]model.Schedule, error)

	//GetScheduleByID returns the Schedule with the given ID
	GetScheduleByID(id int64) (*model.Schedule, error)

	//Insert inserts the Schedule into the database, and returns
	//the newly-inserted ScheduleID, complete with the DBMS-assigned ID
	CreateSchedule(Schedule *model.Schedule) (int64, error)

	//Vote increase the votes for a Schedule, and returns the updated votes
	Vote(id int64) (int, error)

	//Delete deletes the Schedule with the given ID
	DeleteSchedule(id int64) error
}

func dbError(msg string, err error) error {
	if err != nil {
		fmErr := fmt.Errorf("Error when executing [%s] in database: %v", msg, err)
		log.Println(fmErr.Error())
		return fmErr
	}
	return nil
}
