package model

import "time"

//Group struct contains fileds of a group of users

type Group struct {
	GroupID     int64     `json:"groupID"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	Creator     int64     `json:"creator"`
	CreateDate  time.Time `json:"createDate"`
}
