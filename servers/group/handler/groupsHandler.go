package handler

import (
	"MeetingScheduler/servers/group/model"
	"fmt"
	"net/http"
	"path"
)

var errUnsuportMethod = "Unsupported Methods"
var typeText = "text/plain"

// GroupsHandler handles request for creating a group
func (ctx *Context) GroupsHandler(w http.ResponseWriter, r *http.Request) {

	uid := getCurrentUser(w, r)
	if uid < 0 {
		return
	}

	if r.Method != "POST" && r.Method != "GET" {
		http.Error(w, errUnsuportMethod, http.StatusMethodNotAllowed)
		return
	}

	if r.Method == "POST" {
		// Only support JSON body
		if !isContentTypeJSON(w, r) {
			return
		}

		// Read response body
		body := getRequestBody(w, r)
		if body == nil {
			return
		}

		// Unmarshal body to group object
		postedGroup := &model.Group{}
		if !unmarshalBody(w, body, postedGroup) {
			return
		}
		postedGroup.CreatorID = uid

		// Add to database
		id, err := ctx.Store.InsertGroup(postedGroup)
		if !dbErrorHandle(w, "Insert group", err) {
			return
		}

		// Response: TBD
		respMsg := fmt.Sprintf("successfully create group, id: %d\n", id)
		respondWithHeader(w, typeJSON, []byte(respMsg), http.StatusCreated)
	}

	if r.Method == "GET" {
		group, err := ctx.Store.GetAllGroups()
		if !dbErrorHandle(w, "Get all groups", err) {
			return
		}

		res := marshalRep(w, group)
		if res == nil {
			return
		}

		respondWithHeader(w, typeJSON, res, http.StatusCreated)
	}

}

// SpecificGroupsHandler handles request for a specific group with given id. Only creator can
// manage the information of this group. Other users in the group can view the member, meetings
func (ctx *Context) SpecificGroupsHandler(w http.ResponseWriter, r *http.Request) {

	uid := getCurrentUser(w, r)
	if uid < 0 {
		return
	}

	// Only support GET PATCH DELETE method
	if r.Method != "GET" && r.Method != "PATCH" && r.Method != "DELETE" {
		http.Error(w, errUnsuportMethod, http.StatusMethodNotAllowed)
		return
	}

	// parse group id
	urlID := path.Base(r.URL.Path)
	gid := getIDfromURL(w, r, urlID)
	if gid < 0 {
		return
	}

	// read group data from db
	group, err := ctx.Store.GetGroupByID(gid)
	if !dbErrorHandle(w, "Get group", err) {
		return
	}

	// GET request will return the group
	if r.Method == "GET" {

		//Get all meetings
		meetings, err := ctx.Store.GetAllMeetingsOfGroup(gid)
		if !dbErrorHandle(w, "Get all meetings", err) {
			return
		}

		//Get all members of the group
		members, err := ctx.Store.GetAllMembers(gid)
		if !dbErrorHandle(w, "Get all members", err) {
			return
		}

		//Construct the return struct
		completeGroupInfo := model.GroupReturnBody{
			Meetings:  meetings,
			GroupInfo: group,
			Members:   members,
		}

		// marshal current group into response body
		response := marshalRep(w, completeGroupInfo)
		if response == nil {
			return
		}

		// write into response
		respondWithHeader(w, typeJSON, response, http.StatusOK)
	}

	// PATCH request can update the groups' information, only creator can use this method
	if r.Method == "PATCH" {

		// Check authorization
		if !isGroupCreator(group, uid, w) {
			return
		}

		// Check content type
		if !isContentTypeJSON(w, r) {
			return
		}

		// Get request body
		body := getRequestBody(w, r)
		if body == nil {
			return
		}

		// Marshal the body to json
		newGroup := &model.Group{}
		if !unmarshalBody(w, body, newGroup) {
			return
		}

		// Update information in database
		err := ctx.Store.UpdateGroup(newGroup)
		if !dbErrorHandle(w, "Update group", err) {
			return
		}

		// TBD: get the newly updated group
		group, err = ctx.Store.GetGroupByID(group.GroupID)
		if !dbErrorHandle(w, "Get updated group", err) {
			return
		}

		// marshal into body and response
		res := marshalRep(w, group)
		if res == nil {
			return
		}

		respondWithHeader(w, typeJSON, res, http.StatusOK)

	}

	// Delete the current group, only creator can use this method
	if r.Method == "DELETE" {

		// Check authorization
		if !isGroupCreator(group, uid, w) {
			return
		}

		err := ctx.Store.DeleteGroup(gid)
		if !dbErrorHandle(w, "Delete group", err) {
			return
		}

		respondWithHeader(w, typeText, []byte("Delete success"), http.StatusOK)

	}

}

// GroupsMeetingHandler handles request for managing a meeting. A user can create a meeting
// for the group. The creator of a meeting or the creator of a group can modify the meeting
// information
func (ctx *Context) GroupsMeetingHandler(w http.ResponseWriter, r *http.Request) {

	uid := getCurrentUser(w, r)
	if uid < 0 {
		return
	}

	// Only support GET POST method
	if r.Method != "GET" && r.Method != "POST" {
		http.Error(w, errUnsuportMethod, http.StatusMethodNotAllowed)
		return
	}

	// parse group id
	urlID := path.Base(path.Dir(r.URL.Path))
	gid := getIDfromURL(w, r, urlID)
	if gid < 0 {
		return
	}

	// GET method returns all the meetings
	if r.Method == "GET" {
		meetings, err := ctx.Store.GetAllMeetingsOfGroup(gid)
		if !dbErrorHandle(w, "Get all meetings", err) {
			return
		}

		// encode into JSON
		res := marshalRep(w, meetings)
		if res == nil {
			return
		}

		// response
		respondWithHeader(w, typeJSON, res, http.StatusOK)
	}

	// POST method create new meeting
	if r.Method == "POST" {

		// Check content type
		if !isContentTypeJSON(w, r) {
			return
		}

		// Get the post body
		body := getRequestBody(w, r)
		if body == nil {
			return
		}

		// TBD: what fields are available
		meeting := &model.Meeting{}
		if !unmarshalBody(w, body, meeting) {
			return
		}
		meeting.CreatorID = uid
		meeting.GroupID = gid

		// Add into database
		mid, err := ctx.Store.InsertMeeting(meeting)
		if !dbErrorHandle(w, "Insert meeting", err) {
			return
		}

		// TBD: should we return the whole object
		meeting, err = ctx.Store.GetMeetingByID(mid)
		if !dbErrorHandle(w, "get meeting", err) {
			return
		}

		res := marshalRep(w, meeting)
		if res == nil {
			return
		}

		// Response
		respondWithHeader(w, typeText, res, http.StatusCreated)
	}

}

// SpecificGroupsMeetingHandler handles request for a specific group meeting with given id.
func (ctx *Context) SpecificGroupsMeetingHandler(w http.ResponseWriter, r *http.Request) {

	uid := getCurrentUser(w, r)
	if uid < 0 {
		return
	}

	if r.Method != "GET" && r.Method != "PATCH" && r.Method != "DELETE" {
		http.Error(w, errUnsuportMethod, http.StatusMethodNotAllowed)
		return
	}

	// parse group id
	urlID := path.Base(path.Dir(path.Dir(r.URL.Path)))
	gid := getIDfromURL(w, r, urlID)
	if gid < 0 {
		return
	}

	// parse meeting id
	urlID = path.Base(r.URL.Path)
	mid := getIDfromURL(w, r, urlID)
	if mid < 0 {
		return
	}

	if r.Method == "GET" {
		meeting, err := ctx.Store.GetMeetingByID(mid)
		if !dbErrorHandle(w, "Get meeting", err) {
			return
		}

		//get the members of this meeting
		members, err := ctx.Store.GetAllParticipants(mid)
		if !dbErrorHandle(w, "Get all members", err) {
			return
		}

		// get the schedule of the meeting
		schedules, err := ctx.Store.GetAllSchedule(mid)
		if !dbErrorHandle(w, "Get all schedules", err) {
			return
		}

		// construct response body
		completeMeetingInfo := model.MeetingReturnBody{
			MeetingInfo:  meeting,
			Participants: members,
			Schedules:    schedules,
		}

		// encode into JSON
		res := marshalRep(w, completeMeetingInfo)
		if res == nil {
			return
		}

		// response
		respondWithHeader(w, typeJSON, res, http.StatusOK)
	}

	if r.Method == "PATCH" {

		//get object from body
		body := getRequestBody(w, r)
		if body == nil {
			return
		}

		meeting := &model.Meeting{}
		if !unmarshalBody(w, body, meeting) {
			return
		}

		//Update in db
		err := ctx.Store.UpdateMeeting(mid, meeting)
		if !dbErrorHandle(w, "Update meeting", err) {
			return
		}

		// TBD
		meeting, err = ctx.Store.GetMeetingByID(mid)
		if !dbErrorHandle(w, "get updated meeting", err) {
			return
		}

		res := marshalRep(w, meeting)
		if res == nil {
			return
		}

		respondWithHeader(w, typeJSON, res, http.StatusOK)
	}

	if r.Method == "DELETE" {

		// Check authorization
		group, err := ctx.Store.GetGroupByID(gid)
		if !dbErrorHandle(w, "get group", err) || !isGroupCreator(group, uid, w) {
			return
		}

		meeting, err := ctx.Store.GetMeetingByID(mid)
		if !dbErrorHandle(w, "get meeting", err) || !isMeetingCreator(meeting, uid, w) {
			return
		}

		// delete in database
		err = ctx.Store.DeleteMeeting(mid)
		if !dbErrorHandle(w, "delete meeting", err) {
			return
		}

		// TODO: delete associated meeting?

		//response
		respondWithHeader(w, typeText, []byte("Successfully deleted"), http.StatusOK)
	}
}

// ScheduleHandler add, get and delte schedules
func (ctx *Context) ScheduleHandler(w http.ResponseWriter, r *http.Request) {
	// Only support GET POST method
	if r.Method != "GET" && r.Method != "POST" && r.Method != "DELETE" {
		http.Error(w, errUnsuportMethod, http.StatusMethodNotAllowed)
		return
	}

	// parse meeting id
	urlID := path.Base(path.Dir(r.URL.Path))
	mid := getIDfromURL(w, r, urlID)
	if mid < 0 {
		return
	}

	if r.Method == "POST" {
		// Only support JSON body
		if !isContentTypeJSON(w, r) {
			return
		}

		// Read response body
		body := getRequestBody(w, r)
		if body == nil {
			return
		}

		// Unmarshal body to group object
		sch := model.Schedule{}
		if !unmarshalBody(w, body, sch) {
			return
		}
		sch.MeetingID = mid

		// Add to database
		id, err := ctx.Store.CreateSchedule(&sch)
		if !dbErrorHandle(w, "Insert group", err) {
			return
		}

		schedule, err := ctx.Store.GetScheduleByID(id)
		if !dbErrorHandle(w, "Get schedule", err) {
			return
		}

		// marshal into bytes
		response := marshalRep(w, schedule)
		if response == nil {
			return
		}

		// Response
		respondWithHeader(w, typeJSON, response, http.StatusCreated)
	}

	if r.Method == "GET" {
		schedules, err := ctx.Store.GetAllSchedule(mid)
		if !dbErrorHandle(w, "Get schedule", err) {
			return
		}

		response := marshalRep(w, schedules)
		respondWithHeader(w, typeJSON, response, http.StatusOK)
	}

}

func (ctx *Context) SpecificScheduleHandler(w http.ResponseWriter, r *http.Request) {
	// Only support DELETE, PATCH method
	if r.Method != "DELETE" && r.Method != "PATCH" {
		http.Error(w, errUnsuportMethod, http.StatusMethodNotAllowed)
		return
	}

	// parse schedule id
	urlID := path.Base(r.URL.Path)
	sid := getIDfromURL(w, r, urlID)
	if sid < 0 {
		return
	}

	if r.Method == "DELETE" {
		err := ctx.Store.DeleteSchedule(sid)
		if !dbErrorHandle(w, "Delete schedule", err) {
			return
		}

		respondWithHeader(w, typeText, []byte("Successfully deleted the schedule"), http.StatusOK)
	}

	// This method is used to confirm a meeting
	if r.Method == "PATCH" {
		// parse meeting id
		urlID = path.Base(path.Dir(path.Dir(r.URL.Path)))
		mid := getIDfromURL(w, r, urlID)
		if mid < 0 {
			return
		}

		//get schedule
		schedule, err := ctx.Store.GetScheduleByID(sid)
		if !dbErrorHandle(w, "Get schedule", err) {
			return
		}

		// Confim current meeting with current schedule
		meeting, err := ctx.Store.ConfirmMeeting(mid, schedule)

		// marshal response
		response := marshalRep(w, meeting)

		respondWithHeader(w, typeJSON, response, http.StatusOK)

	}

}
