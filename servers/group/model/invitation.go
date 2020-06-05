package model

import (
	"fmt"
	"math"
	"math/rand"
)

// Guest stores information for non-registered users
type Guest struct {
	GuestID     int64  `json:"guestID"`
	Email       string `json:"email"`
	DisplayName string `json:"name"`
	GroupID     int64  `json:"groupID"`
	MeetingID   int64  `json:"meetingID"`
	InvitedBy   int64  `json:"invitedBy"`
	Confirmed   int    `json:"confirmed"`
}

// NewGuest is used for accept request body
type NewGuest struct {
	Email       string `json:"email"`
	DisplayName string `json:"name"`
}

// DisplayGuest is used to display for the clients
type DisplayGuest struct {
	Email       string `json:"email"`
	DisplayName string `json:"name"`
	Confirmed   int    `json:"confirmed"`
}

// Invitation stores the meeting, guest information
// It's used to response to the guest page
type Invitation struct {
	MeetingName   string      `json:"meetingName"`
	MeetingDetail string      `json:"meetingDetail"`
	GuestName     string      `json:"guestName"`
	GuestEmail    string      `json:"guestEmail"`
	InvitedBy     string      `json:"invitedBy"`
	CreateDate    string      `json:"createDate"`
	Schedules     []*Schedule `json:"meetingSchedule"`
}

// CreateGuest creates a new guest with the provided email and name
func CreateGuest(guest Guest, gid int64, mid int64, uid int64) *Guest {
	randomID := int64(rand.Float64() * math.Pow(10, 7))

	return &Guest{
		GuestID:     randomID,
		Email:       guest.Email,
		DisplayName: guest.DisplayName,
		GroupID:     gid,
		MeetingID:   mid,
		InvitedBy:   uid,
		Confirmed:   0,
	}
}

func GenerateRandomID() int64 {
	return int64(rand.Float64() * math.Pow(10, 7))
}

// CreateMeetingInvitation generates a random link for guest
func CreateMeetingInvitation(guestName string, email string, mid int64) string {
	return fmt.Sprintf("m=%d&name=%s&email=%s&id=%d", mid, guestName, email, GenerateRandomID())
}

// CreateGroupInvitation generates a random link for guest
func CreateGroupInvitation(guestName string, email string, gid int64) string {
	return fmt.Sprintf("g=%d&name=%s&email=%s&id=%d", gid, guestName, email, GenerateRandomID())
}
