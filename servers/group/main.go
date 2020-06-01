package main

import (
	"log"
	"net/http"
	"os"

	h "MeetingScheduler/servers/group/handler"

	"github.com/gorilla/mux"
)

func main() {
	ADDR := os.Getenv("ADDR")

	//context
	ctx := h.Context{}

	// handlers
	route := mux.NewRouter()
	route.HandleFunc("/v1/groups", ctx.GroupsHandler)
	route.HandleFunc("/v1/groups/", ctx.SpecificGroupsHandler)
	route.HandleFunc("/v1/groups/{group_id}/meetings", ctx.GroupsMeetingHandler)
	route.HandleFunc("/v1/groups/{group_id}/meetings/{meeting_id}", ctx.SpecificGroupsMeetingHandler)

	// start server
	log.Printf("Server is listening at %s...", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, route))
}
