package dao

import "MeetingScheduler/servers/group/model"

var getAllGuestByMeeting = "SELECT Email, DisplayName, Confirmed FROM guests WHERE MeetingID = ?"
var getAllGuestByGroup = "SELECT Email, DisplayName, Confirmed FROM guests WHERE GroupID = ?"

// For testing
var insertGuest = "INSERT INTO guests(GuestID, Email, displayName, GroupID, MeetingID, InvitedBy, Confirmed) VALUES(?,?,?,?,?,?,?)"

// GetAllGuestOfMeeting get all guests of the given meeting id
func (store *Store) GetAllGuestOfMeeting(mid int64) ([]*model.DisplayGuest, error) {
	var guests []*model.DisplayGuest
	//uid, email, userName, firstName, lastName
	rows, err := store.Db.Query(getAllGuestByMeeting, mid)
	if err != nil {
		return guests, err
	}

	defer rows.Close()

	for rows.Next() {
		var g model.DisplayGuest
		err = rows.Scan(&g.Email, &g.DisplayName, &g.Confirmed)
		if err != nil {
			return guests, err
		}
		guests = append(guests, &g)
	}

	// get any error encountered during iteration
	return guests, rows.Err()
}

// GetAllGuestOfGroup get the guests of certain group
func (store *Store) GetAllGuestOfGroup(gid int64) ([]*model.DisplayGuest, error) {
	var guests []*model.DisplayGuest
	//uid, email, userName, firstName, lastName
	rows, err := store.Db.Query(getAllGuestByGroup, gid)
	if err != nil {
		return guests, err
	}

	defer rows.Close()

	for rows.Next() {
		var g model.DisplayGuest
		err = rows.Scan(&g.Email, &g.DisplayName, &g.Confirmed)
		if err != nil {
			return guests, err
		}
		guests = append(guests, &g)
	}

	// get any error encountered during iteration
	return guests, rows.Err()
}

// InsertGuest inserts a guest into db
func (store *Store) InsertGuest(id int64, email string, name string, gid int64, mid int64, uid int64) (int64, error) {

	_, err := store.Db.Exec(insertGuest, id, email, name, gid, mid, uid, 0)
	if err != nil {
		return 0, err
	}

	return id, nil
}
