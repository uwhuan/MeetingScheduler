package model

import "time"

//Group struct contains fileds of a group of users

type Group struct {
	GroupID    int       `json:"groupID"`
	Name       string    `json:"name"`
	Creator    string    `json:"creator"`
	CreateDate time.Time `json:"createDate"`
}
