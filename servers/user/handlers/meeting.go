package handlers

import (
	"MeetingScheduler/servers/user/model"
	"MeetingScheduler/servers/user/sessions"
	"encoding/json"
	"log"
	"net/http"
	"path"
	"reflect"
	"strconv"
	"time"
)

func (ctx *HandlerCtx) GetUserIDFromSession(r *http.Request) (int64, error) {
	session, err := sessions.GetSessionID(r, ctx.SigningKey)
	if err != nil {
		return -1, err
	}
	stateRet := SessionState{}
	err = ctx.SessionStore.Get(session, &stateRet)
	if err != nil {
		log.Println("error getting userID from session", err)
		return -1, err
	}

	return stateRet.User.UID, nil
}

/**
itemExists is used to validate parameters
*/
func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

/**
filterMeetings is used to filter the meetings based off the expected time parameter
*/
func filterMeetings(meetings []*model.Meeting, timeParam string) []*model.Meeting {
	if timeParam == "all" {
		return meetings
	}
	returnValue := []*model.Meeting{}

	for i := 0; i < len(meetings); i++ {
		currentMeeting := meetings[i]
		if timeParam == "past" && time.Now().After(currentMeeting.StartTime) {
			returnValue = append(returnValue, currentMeeting)
		} else if timeParam == "future" && time.Now().Before(currentMeeting.StartTime) {
			returnValue = append(returnValue, currentMeeting)
		}
	}

	return returnValue
}

/**
GetUserMeetingsHandler handles requests for the "meetings" resource.
GET request meetings that a user has created/have attended/will attended
*/
func (ctx *HandlerCtx) GetUserMeetingsHandler(w http.ResponseWriter, r *http.Request) {

	userId, err := ctx.GetUserIDFromSession(r)
	if err != nil {
		http.Error(w, "Login required for request", http.StatusUnauthorized)
		return
	}

	timeParam := r.URL.Query().Get("time")
	timeArray := [3]string{"all", "past", "future"}
	if len(timeParam) == 0 || !itemExists(timeArray, timeParam) {
		http.Error(w, "time param not supplied", http.StatusBadRequest)
		return
	}

	viewtype := r.URL.Query().Get("viewtype")
	viewtypeArray := [2]string{"all", "created"}
	if len(viewtype) == 0 || !itemExists(viewtypeArray, viewtype) {
		http.Error(w, "viewtype param not supplied", http.StatusBadRequest)
		return
	}

	meetings := []*model.Meeting{}

	if viewtype == "all" {
		meetings, err = ctx.UserStore.GetAllUserMeetings(userId)
	} else {
		meetings, err = ctx.UserStore.GetMeetingsByCreatorID(userId)
	}

	if err != nil {
		http.Error(w, "Error fetching meetings from database", http.StatusInternalServerError)
		return
	}

	filteredMeetings := filterMeetings(meetings, timeParam)

	// Respond to client
	response, err := json.Marshal(filteredMeetings)
	if err != nil {
		http.Error(w, "Error marshalling meetings", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

/**
GetMeetingByIDHandler handles requests for the "meetings" resource.
GET request getting a meeting by id
*/
func (ctx *HandlerCtx) GetMeetingByIDHandler(w http.ResponseWriter, r *http.Request) {

	_, err := ctx.GetUserIDFromSession(r)
	if err != nil {
		http.Error(w, "Login required for request", http.StatusUnauthorized)
		return
	}

	meetingId := path.Base(r.URL.Path)
	meetingIdInt, _ := strconv.ParseInt(meetingId, 10, 64)
	meeting, err := ctx.UserStore.GetMeetingByID(meetingIdInt)

	if err != nil {
		http.Error(w, "Error getting the meeting from database", http.StatusInternalServerError)
		return
	} else if meeting == nil {
		http.Error(w, "Meeting not found", http.StatusNotFound)
		return
	}

	participantIds, err := ctx.UserStore.GetMeetingParticipants(meetingIdInt)
	if err != nil {
		http.Error(w, "Error getting the participants from database", http.StatusInternalServerError)
		return
	}

	meetingDetails := model.MeetingDetails{
		Meeting:      *meeting,
		Participants: participantIds,
	}

	// Respond to client
	response, err := json.Marshal(meetingDetails)
	if err != nil {
		http.Error(w, "Error marshalling meeting", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
