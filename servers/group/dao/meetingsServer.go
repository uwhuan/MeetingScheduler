package dao

import (
	model "MeetingScheduler/servers/group/model"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// MeetingDB is a struct of sql dababase
type MeetingDB struct {
	Db *sql.DB
}

var queryGetMeeting = "SELECT meetingID, name, description, creator, startTime, endTime, createDate, confirmed, groupID FROM meetings where meetingID = ?"
var queryGetAllByGroup = "SELECT name, creator, startTime, endTime, createDate, confirmed, meetingID, groupID FROM meetings where groupID = ?"
var queryInsertMeeting = "INSERT INTO meetings(name, description, creator, startTime, endTime, createDate, confirmed, groupID) VALUES (?,?,?,?,?,?,?,?)"
var queryUpdateMeeting = "UPDATE meetings SET name = ?, description = ? WHERE id = ?"
var queryDeleteMeeting = "DELETE FROM meetings WHERE id = ?"
var queryConfirmMeeting = "UPDATE meetings SET confirmed = true, startTime = ?, endTime = ? WHERE id = ?"

var queryGetAllParticipants = "SELECT uid, email, userName, firstName, lastName FROM user INNER JOIN meetingparticipant M ON M.uid =user.uid WHERE M.meetingID = ?"

var defaultTime = "Jan 1, 2000 at 0:00pm (PST)"
var defaultConfrim = false

// var defaultErrorMsg = "handle meetings"

//GetMeetingByID returns the meeting with the given ID
func (store *MeetingDB) GetMeetingByID(id int64) (*model.Meeting, error) {

	var meeting model.Meeting
	// Execute the query
	err := store.Db.QueryRow(queryGetMeeting, id).Scan(&meeting.MeetingID, &meeting.Name,
		&meeting.Description, &meeting.Creator, &meeting.StartTime, &meeting.EndTime,
		&meeting.CreateDate, &meeting.Confirmed, &meeting.GroupID)
	return &meeting, err

}

//GetAllMeetingsOfGroup returns all the meeting with the given groupID
func (store *MeetingDB) GetAllMeetingsOfGroup(id int64) ([]*model.Meeting, error) {
	var meetings []*model.Meeting
	rows, err := store.Db.Query(queryGetAllByGroup, id)
	if err != nil {
		return meetings, err
	}

	defer rows.Close()

	for rows.Next() {
		var meeting model.Meeting
		err = rows.Scan(&meeting.Name, &meeting.Description, &meeting.Creator,
			&meeting.StartTime, &meeting.EndTime, &meeting.CreateDate, &meeting.Confirmed,
			&meeting.MeetingID, &meeting.GroupID)
		if err != nil {
			return meetings, err
		}
		meetings = append(meetings, &meeting)
	}

	// get any error encountered during iteration
	return meetings, rows.Err()
}

//InsertMeeting inserts the meeting into the database, and returns
//the newly-inserted meetingID, complete with the DBMS-assigned ID
func (store *MeetingDB) InsertMeeting(meeting *model.Meeting) (int64, error) {

	// Execute the query
	res, err := store.Db.Exec(queryInsertMeeting, meeting.Name, meeting.Description, meeting.Creator,
		parseTime(defaultTime), parseTime(defaultTime), time.Now, defaultConfrim, meeting.GroupID)
	if err != nil {
		return 0, err
	}

	// Get the auto-incremented id
	id, err := res.LastInsertId()
	return id, err

}

//UpdateMeetingName applies updates to the given meeting  ID
//and returns any errors
func (store *MeetingDB) UpdateMeetingName(id int64, update *model.Meeting) error {
	_, err := store.Db.Exec(queryUpdateMeeting, update.Name, update.Description, id)
	return err
}

//DeleteMeeting deletes the meeting with the given ID
func (store *MeetingDB) DeleteMeeting(id int64) error {
	_, err := store.Db.Exec(queryDeleteMeeting, id)
	return err
}

// ConfirmMeeting set the confirmed start and end time of a meeting
// and set the confirmed flag to be true
func (store *MeetingDB) ConfirmMeeting(id int64, schedule *model.Schedule) error {
	_, err := store.Db.Exec(queryConfirmMeeting, parseTime(schedule.StartTime), parseTime(schedule.EndTime), id)
	return err
}

//GetAllParticipants get all participants of a meeting
func (store *MeetingDB) GetAllParticipants(meetingID int64) ([]*model.User, error) {
	var users []*model.User
	//uid, email, userName, firstName, lastName
	rows, err := store.Db.Query(queryGetAllParticipants, meetingID)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var u *model.User
		err = rows.Scan(u.ID, u.Email, u.UserName, u.FirstName, u.LastName)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}

	// get any error encountered during iteration
	return users, rows.Err()
}

// parseTime is a helper function to parse a string to time format
// the input string should follow this format: "Jan 1, 2000 at 0:00pm (PST)"
func parseTime(timeStr string) time.Time {
	t, err := time.Parse(defaultTime, timeStr)
	if err != nil {
		fmErr := fmt.Errorf("Error when trying to parse [%s]: %v", timeStr, err)
		log.Println(fmErr.Error())
	}
	return t
}
