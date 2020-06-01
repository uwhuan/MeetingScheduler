package dao

import (
	model "MeetingScheduler/servers/group/model"
)

var queryGetAllSchedule = "SELECT StartTime, EndTime, votes FROM schedule WHERE MeetingID = ? ORDER BY votes"
var queryInsertSchedule = "INSERT INTO schedule(meetingID, StartTime, EndTime, votes) VALUES (?,?,?,?)"
var queryGetSchedule = "SELECT scheduleID, StartTime, EndTime, MeetingID, votes FROM schedule WHERE scheduleID = ?"
var queryDeleteSchedule = "DELETE FROM schedule WHERE scheduleID = ?"
var queryIncreaseVote = "UPDATE schedule SET votes = votes+1 WHERE scheduleID = ?"
var queryCheckVotes = "SELECT votes FROM schedule WHERE scheduleID = ?"

//GetAllSchedule returns all schedules unser a meeting
func (store *Store) GetAllSchedule(meetingID int64) ([]*model.Schedule, error) {

	var schedules []*model.Schedule
	//uid, email, userName, firstName, lastName
	rows, err := store.Db.Query(queryGetAllSchedule, meetingID)
	if err != nil {
		return schedules, err
	}

	defer rows.Close()

	for rows.Next() {
		var sch *model.Schedule
		err = rows.Scan(sch.StartTime, sch.EndTime, sch.MeetingID)
		if err != nil {
			return schedules, err
		}
		schedules = append(schedules, sch)
	}

	// get any error encountered during iteration
	return schedules, rows.Err()
}

//GetScheduleByID returns the Schedule with the given ID
func (store *Store) GetScheduleByID(id int64) (*model.Schedule, error) {
	var sch model.Schedule
	err := store.Db.QueryRow(queryGetSchedule, id).Scan(&sch.ScheduleID, &sch.StartTime, &sch.EndTime, &sch.MeetingID, &sch.Votes)
	return &sch, err
}

//CreateSchedule the Schedule into the database, and returns
//the newly-inserted ScheduleID, complete with the DBMS-assigned ID
func (store *Store) CreateSchedule(schedule *model.Schedule) (int64, error) {
	res, err := store.Db.Exec(queryInsertSchedule, schedule.MeetingID, schedule.StartTime, schedule.EndTime, schedule.Votes)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

//Vote increase the votes for a Schedule
func (store *Store) Vote(id int64) (int, error) {
	_, err := store.Db.Exec(queryIncreaseVote, id)
	if err != nil {
		return 0, err
	}
	voteCount := 0
	err = store.Db.QueryRow(queryCheckVotes, id).Scan(voteCount)
	return voteCount, err
}

//DeleteSchedule deletes the Schedule with the given ID
func (store *Store) DeleteSchedule(id int64) error {
	_, err := store.Db.Exec(queryDeleteSchedule, id)
	return err
}
