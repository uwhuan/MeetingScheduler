package model

import (
	"errors"
)

//ErrUserNotFound is returned when the user can't be found
var ErrUserNotFound = errors.New("user not found")

//Store represents a store for Users and their meetings
type Store interface {
	//GetByID returns the User with the given ID
	GetByID(id int64) (*User, error)

	//GetByEmail returns the User with the given email
	GetByEmail(email string) (*User, error)

	//GetByUserName returns the User with the given Username
	GetByUserName(username string) (*User, error)

	//Insert inserts the user into the database, and returns
	//the newly-inserted User, complete with the DBMS-assigned ID
	Insert(user *User) (*User, error)

	//InsertLog logs all user sign-in attempts into the database
	InsertLog(logEntry *LogEntry) error

	//Update applies UserUpdates to the given user ID
	//and returns the newly-updated user
	Update(id int64, updates *Updates) (*User, error)

	//Delete deletes the user with the given ID
	Delete(id int64) error

	//GetMeetingByID returns the meeting with the given meeting ID
	GetMeetingByID(id int64) (*Meeting, error)

	//GetMeetingsByCreatorID returns the meeting that the user has created
	GetMeetingsByCreatorID(id int64) ([]*Meeting, error)

	//GetAllUserMeetings returns all of the user's meetings
	GetAllUserMeetings(id int64) ([]*Meeting, error)

	//GetMeetingParticipants returns participants for meeting
	GetMeetingParticipants(id int64) ([]int64, error)
}
