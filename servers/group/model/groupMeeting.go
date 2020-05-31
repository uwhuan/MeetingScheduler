package model

type groupMeeting struct {
	groupID   int
	meetingID int
	gmID      int
}

func CreateGroupMeeting(gm *groupMeeting) error {
	return nil
}

func GetGroupMeeting() *groupMeeting {
	return nil
}

func DeleteGroupMeeting(gmID int) error {
	return nil
}
