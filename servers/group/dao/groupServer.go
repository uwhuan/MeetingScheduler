package dao

import (
	model "MeetingScheduler/servers/group/model"
	"time"
)

var queryGetGroup = "SELECT groupID, name, description, creatorID, createDate FROM userGroups where groupID = ?"
var queryInsertGroup = "INSERT INTO userGroups(name, description, creatorID, createDate) VALUES (?,?,?,?)"
var queryUpdateGroup = "UPDATE userGroups SET name = ?, description = ? WHERE groupID = ?"
var queryDeleteGroup = "DELETE FROM userGroups WHERE groupID = ?"
var queryGetAllGroups = "SELECT groupID, name, description, creatorID, createDate FROM userGroups"

var queryGetAllMembers = "SELECT user.uid, email, userName, firstName, lastName FROM user INNER JOIN membership M ON user.uid = M.uid WHERE M.groupID = ?"

var queryInsertGU = "INSERT INTO membership(GroupID, uid) VALUES(?,?)"
var queryDeleteMeetingsOfGroup = "DELETE FROM meetings WHERE groupID = ?"
var queryDeleteParticipantsOfGroup = "DELETE FROM membership WHERE groupID = ?"

//GetGroupByID returns the Group with the given ID
func (store *Store) GetGroupByID(id int64) (*model.Group, error) {
	var group model.Group
	err := store.Db.QueryRow(queryGetGroup, id).Scan(&group.GroupID, &group.Name, &group.Description, &group.CreatorID, &group.CreateDate)
	return &group, err
}

//InsertGroup inserts the Group into the database, and returns
//the newly-inserted GroupID, complete with the DBMS-assigned ID
func (store *Store) InsertGroup(group *model.Group) (int64, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0, err
	}

	res, err := store.Db.Exec(queryInsertGroup, group.Name, group.Description, group.CreatorID, time.Now().In(loc).Format(time.UnixDate))
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	_, err = store.Db.Exec(queryInsertGU, id, group.CreatorID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

//UpdateGroup applies updates to the given Group  ID
//and returns the id of the newly-updated Group
func (store *Store) UpdateGroup(update *model.Group) error {
	_, err := store.Db.Exec(queryUpdateGroup, update.Name, update.Description, update.GroupID)
	return err
}

//DeleteGroup deletes the Group and associted meetings with the given ID
func (store *Store) DeleteGroup(id int64) error {
	_, err := store.Db.Exec(queryDeleteGroup, id)
	if err != nil {
		return err
	}

	//delete all meetings of the group
	_, err = store.Db.Exec(queryDeleteMeetingsOfGroup, id)

	//delete all members of the group
	_, err = store.Db.Exec(queryDeleteParticipantsOfGroup, id)
	return err
}

//GetAllMembers gets all users in the current group
func (store *Store) GetAllMembers(id int64) ([]*model.User, error) {
	var users []*model.User
	//uid, email, userName, firstName, lastName
	rows, err := store.Db.Query(queryGetAllMembers, id)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var u model.User
		err = rows.Scan(&u.ID, &u.Email, &u.UserName, &u.FirstName, &u.LastName)
		if err != nil {
			return users, err
		}
		users = append(users, &u)
	}

	// get any error encountered during iteration
	return users, rows.Err()
}

//GetAllGroups gets all groups in db
func (store *Store) GetAllGroups() ([]*model.Group, error) {
	var groups []*model.Group
	//uid, email, userName, firstName, lastName
	rows, err := store.Db.Query(queryGetAllGroups)
	if err != nil {
		return groups, err
	}

	defer rows.Close()

	for rows.Next() {
		var g model.Group
		err = rows.Scan(&g.GroupID, &g.Name, &g.Description, &g.CreatorID, &g.CreateDate)
		if err != nil {
			return groups, err
		}
		groups = append(groups, &g)
	}

	// get any error encountered during iteration
	return groups, rows.Err()
}
