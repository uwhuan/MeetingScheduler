package handler

import (
	"MeetingScheduler/servers/group/model"
	"fmt"
	"net/http"
)

var errUnsuportMethod = "Unsupported Methods"
var typeText = "text/plain"

// GroupsHandler handles request for creating a group
func (ctx *Context) GroupsHandler(w http.ResponseWriter, r *http.Request) {

	uid := getCurrentUser(w, r)
	if uid < 0 {
		return
	}

	// Only support POST method
	if r.Method != "POST" {
		http.Error(w, errUnsuportMethod, http.StatusMethodNotAllowed)
		return
	}

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

	//TBD: add any validation?

	// Add to database
	id, err := ctx.Store.InsertGroup(postedGroup)
	if dbErrorHandle(w, "Insert group", err) {
		return
	}

	// Response: TBD
	respMsg := fmt.Sprintf("successfully create group, id: %d\n", id)
	respondWithHeader(w, typeJSON, []byte(respMsg), http.StatusCreated)

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
	id := getIDfromURL(w, r)
	if id < 0 {
		return
	}

	// read group data from db
	group, err := ctx.Store.GetGroupByID(id)
	if !dbErrorHandle(w, "Get group", err) {
		return
	}

	// GET request will return the group
	if r.Method == "GET" {

		// marshal current group into response body
		response := marshalRep(w, group)
		if response == nil {
			return
		}

		//TODO: also add all meetings, members associated

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

		err := ctx.Store.DeleteGroup(id)
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
	id := getGroupIDfromURL(w, r)
	if id < 0 {
		return
	}

	// GET method returns all the meetings
	if r.Method == "GET" {
		meetings, err := ctx.Store.GetAllMeetingsOfGroup(id)
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

		// Add into database
		id, err := ctx.Store.InsertMeeting(meeting)
		if !dbErrorHandle(w, "Insert meeting", err) {
			return
		}

		// TBD: should we return the whole object
		meeting.MeetingID = id
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

	// parse meeting id
	id := getIDfromURL(w, r)
	if id < 0 {
		return
	}

	if r.Method == "GET" {
		meetings, err := ctx.Store.GetMeetingByID(id)
		if !dbErrorHandle(w, "Get meetings", err) {
			return
		}

		// encode into JSON
		res := marshalRep(w, meetings)
		if res == nil {
			return
		}

		//TODO: also add all members associated

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
		err := ctx.Store.UpdateMeeting(id, meeting)
		if !dbErrorHandle(w, "Update meeting", err) {
			return
		}

		// TBD
		meeting, err = ctx.Store.GetMeetingByID(id)
		if dbErrorHandle(w, "get updated meeting", err) {
			return
		}

		res := marshalRep(w, meeting)
		if res == nil {
			return
		}

		respondWithHeader(w, typeJSON, res, http.StatusOK)
	}

	if r.Method == "DELETE" {

		// parse group id
		gid := getGroupIDfromURL(w, r)
		if gid < 0 {
			return
		}

		// Check authorization
		group, err := ctx.Store.GetGroupByID(gid)
		if !dbErrorHandle(w, "get group", err) || !isGroupCreator(group, uid, w) {
			return
		}

		meeting, err := ctx.Store.GetMeetingByID(id)
		if !dbErrorHandle(w, "get meeting", err) || !isMeetingCreator(meeting, uid, w) {
			return
		}

		// delete in database
		err = ctx.Store.DeleteMeeting(id)
		if !dbErrorHandle(w, "delete meeting", err) {
			return
		}

		//response
		respondWithHeader(w, typeText, []byte("Successfully deleted"), http.StatusOK)
	}
}
