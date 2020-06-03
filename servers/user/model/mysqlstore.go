package model

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlStore struct {
	Db *sql.DB
}

/**
GetByID returns the User with the given ID
*/
func (mysqlStore *MySqlStore) GetByID(id int64) (*User, error) {

	var user User
	// Execute the query
	err := mysqlStore.Db.QueryRow("SELECT UID, Email, PassHash, UserName, FirstName, LastName FROM user where UID = ?", id).Scan(&user.UID, &user.Email, &user.PassHash, &user.UserName, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

/**
GetByEmail returns the User with the given email
*/
func (mysqlStore *MySqlStore) GetByEmail(email string) (*User, error) {

	var user User
	// Execute the query
	err := mysqlStore.Db.QueryRow("SELECT UID, Email, PassHash, UserName, FirstName, LastName FROM user where Email = ?", email).Scan(&user.UID, &user.Email, &user.PassHash, &user.UserName, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

/**
GetByUserName returns the User with the given Username
*/
func (mysqlStore *MySqlStore) GetByUserName(username string) (*User, error) {

	var user User
	// Execute the query
	err := mysqlStore.Db.QueryRow("SELECT UID, Email, PassHash, UserName, FirstName, LastName FROM user where Username = ?", username).Scan(&user.UID, &user.Email, &user.PassHash, &user.UserName, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

/**
Insert inserts the user into the database, and returns
the newly-inserted User, complete with the DBMS-assigned ID
*/
func (mysqlStore *MySqlStore) Insert(user *User) (*User, error) {

	// perform a db.Exec insert
	insert, err := mysqlStore.Db.Exec("INSERT INTO user VALUES (?,?,?,?,?,?)", nil, user.Email, user.PassHash, user.UserName, user.FirstName, user.LastName)

	// if there is an error inserting, handle it
	// if not, returning the insert ID
	if err != nil {
		return nil, err
	}

	id, err := insert.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.UID = id
	return user, nil
}

/**
InsertLog logs all user sign-in attempts into the database
*/
func (mysqlStore *MySqlStore) InsertLog(logEntry *LogEntry) error {
	// perform a db.Exec insert
	_, err := mysqlStore.Db.Exec("INSERT INTO logentry VALUES (?,?,?)", logEntry.UserID, logEntry.SignInTime, logEntry.IPAddress)

	// if there is an error inserting, handle it
	// if not, returning the insert ID
	if err != nil {
		return err
	}

	return nil
}

/**
Update applies UserUpdates to the given user ID
and returns the newly-updated user
*/
func (mysqlStore *MySqlStore) Update(id int64, updates *Updates) (*User, error) {

	// Get the user by given id
	user, err := mysqlStore.GetByID(id)
	if err != nil {

		return nil, err
	}

	// Apply the updates on the user
	err = user.ApplyUpdates(updates)
	if err != nil {
		return nil, err
	}

	// perform a db.Exec updates
	_, err = mysqlStore.Db.Exec("UPDATE user SET FirstName = ?, LastName = ? WHERE uid = ?", user.FirstName, user.LastName, id)

	// if there is an error updating, handle it
	if err != nil {
		return nil, err
	}

	return user, nil
}

/**
Delete deletes the user with the given ID
*/
func (mysqlStore *MySqlStore) Delete(id int64) error {

	// perform a db.Exec delete
	_, err := mysqlStore.Db.Exec("DELETE FROM user WHERE uid = ?", id)

	// if there is an error deleting, handle it
	if err != nil {
		return err
	}

	return nil
}

/**
GetMeetingByID returns the meeting with the given meeting ID
*/
func (mysqlStore *MySqlStore) GetMeetingByID(id int64) (*Meeting, error) {

	var meeting Meeting
	// Execute the query
	err := mysqlStore.Db.QueryRow("SELECT MeetingID, Name, CreatorID, GroupID, StartTime, EndTime, CreateDate, Description, Confirmed FROM meetings where MeetingID = ?", id).Scan(&meeting.MeetingID, &meeting.Name, &meeting.CreatorID, &meeting.GroupID, &meeting.StartTime, &meeting.EndTime, &meeting.CreateDate, &meeting.Description, &meeting.Confirmed)
	if err != nil {
		return nil, err
	}

	return &meeting, nil
}

/**
GetMeetingsByCreatorID returns the meeting that the user has created
*/
func (mysqlStore *MySqlStore) GetMeetingsByCreatorID(id int64) ([]*Meeting, error) {

	returnValue := []*Meeting{}
	rows, err := mysqlStore.Db.Query("SELECT MeetingID, Name, CreatorID, GroupID, StartTime, EndTime, CreateDate, Description, Confirmed FROM meetings where CreatorID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var meeting Meeting
		// Execute the query
		err := rows.Scan(&meeting.MeetingID, &meeting.Name, &meeting.CreatorID, &meeting.GroupID, &meeting.StartTime, &meeting.EndTime, &meeting.CreateDate, &meeting.Description, &meeting.Confirmed)
		if err != nil {
			return nil, err
		}
		returnValue = append(returnValue, &meeting)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return returnValue, nil
}

/**
GetMeetingParticipants returns participants for meeting
*/
func (mysqlStore *MySqlStore) GetMeetingParticipants(id int64) ([]int64, error) {
	userIDs := []int64{}
	rows, err := mysqlStore.Db.Query("SELECT UID FROM meetingparticipant where MeetingID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {

		var userID int64
		err = rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return userIDs, nil
}

/**
GetAllUserMeetings returns all of the user's meetings
*/
func (mysqlStore *MySqlStore) GetAllUserMeetings(id int64) ([]*Meeting, error) {

	meetingIDs := []*int64{}
	rows, err := mysqlStore.Db.Query("SELECT MeetingID FROM meetingparticipant where UID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {

		var meetingID int64
		err = rows.Scan(&meetingID)
		if err != nil {
			return nil, err
		}
		meetingIDs = append(meetingIDs, &meetingID)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	returnValue := []*Meeting{}
	stmt, err := mysqlStore.Db.Prepare("SELECT MeetingID, Name, CreatorID, GroupID, StartTime, EndTime, CreateDate, Description, Confirmed FROM meetings where MeetingID in (?" + strings.Repeat(",?", len(meetingIDs)-1) + ")")
	if err != nil {
		return nil, err
	}

	rows, err = stmt.Query(meetingIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var meeting Meeting
		// Execute the query
		err := rows.Scan(&meeting.MeetingID, &meeting.Name, &meeting.CreatorID, &meeting.GroupID, &meeting.StartTime, &meeting.EndTime, &meeting.CreateDate, &meeting.Description, &meeting.Confirmed)
		if err != nil {
			return nil, err
		}
		returnValue = append(returnValue, &meeting)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return returnValue, nil
}
