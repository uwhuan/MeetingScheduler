package model

//Group struct contains fileds of a group of users
type Group struct {
	GroupID     int64  `json:"groupID"`
	Description string `json:"description"`
	Name        string `json:"name"`
	CreatorID   int64  `json:"creatorID"`
	CreateDate  string `json:"createDate"`
}

// GroupReturnBody associated the group with memebers and meetings
// It is used to return a detailed group information to the clients
type GroupReturnBody struct {
	GroupInfo *Group
	Meetings  []*Meeting
	Members   []*User
}
