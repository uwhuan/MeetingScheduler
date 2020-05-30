package handler

import "net/http"

// GroupsHandler handles request for creating a group
func (ctx *HandlerContext) GroupsHandler(w http.ResponseWriter, r *http.Request) {

}

// SpecificGroupsHandler handles request for a specific group with given id. Only creator can
// manage the information of this group. Other users in the group can view the member, meetings
func (ctx *HandlerContext) SpecificGroupsHandler(w http.ResponseWriter, r *http.Request) {

}

// GroupsMeetingHandler handles request for managing a meeting. A user can create a meeting
// for the group. The creator of a meeting or the creator of a group can modify the meeting
// information
func (ctx *HandlerContext) GroupsMeetingHandler(w http.ResponseWriter, r *http.Request) {

}

// SpecificGroupsMeetingHandler handles request for a specific group meeting with given id.
func (ctx *HandlerContext) SpecificGroupsMeetingHandler(w http.ResponseWriter, r *http.Request) {

}
