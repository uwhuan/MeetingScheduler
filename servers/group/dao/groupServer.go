package dao

import (
	model "MeetingScheduler/servers/group/model"
	"database/sql"
	"time"
)

// GroupDB is a struct of sql dababase
type GroupDB struct {
	Db *sql.DB
}

var queryGetGroup = "SELECT groupID, name, description, creatorID, createDate, FROM groups where groupID = ?"
var queryInsertGroup = "INSERT INTO groups VALUES (?,?,?,?)"
var queryUpdateGroup = "UPDATE groups SET name = ?, description = ? WHERE groupID = ?"
var queryDeleteGroup = "DELETE FROM groups WHERE groupID = ?"

var queryGetAllMembers = "SELECT uid, email, userName, firstName, lastName FROM user INNER JOIN membership M ON user.uid = M.uid WHERE M.groupID = ?"

//GetGroupByID returns the Group with the given ID
func (store *GroupDB) GetGroupByID(id int64) (*model.Group, error) {
	var group *model.Group
	err := store.Db.QueryRow(queryGetGroup, id).Scan(group.GroupID, group.Name, group.Description, group.CreatorID, group.CreateDate)
	return group, err
}

//InsertGroup inserts the Group into the database, and returns
//the newly-inserted GroupID, complete with the DBMS-assigned ID
func (store *GroupDB) InsertGroup(group *model.Group) (int64, error) {
	res, err := store.Db.Exec(queryInsertGroup, group.Name, group.Description, group.CreatorID, time.Now)
	if err != nil {
		return 0, nil
	}

	return res.LastInsertId()
}

//UpdateGroup applies updates to the given Group  ID
//and returns the id of the newly-updated Group
func (store *GroupDB) UpdateGroup(update *model.Group) error {
	_, err := store.Db.Exec(queryUpdateGroup, update.Name, update.Description, update.GroupID)
	return err
}

//DeleteGroup deletes the Group with the given ID
func (store *GroupDB) DeleteGroup(id int64) error {
	_, err := store.Db.Exec(queryDeleteGroup, id)
	return err
}

//GetAllMembers gets all users in the current group
func (store *GroupDB) GetAllMembers(id int64) ([]*model.User, error) {
	var users []*model.User
	//uid, email, userName, firstName, lastName
	rows, err := store.Db.Query(queryGetAllMembers, id)
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
