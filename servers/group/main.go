package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"MeetingScheduler/servers/group/dao"
	h "MeetingScheduler/servers/group/handler"

	"github.com/gorilla/mux"
)

func main() {
	ADDR := os.Getenv("ADDR")
	dsn := os.Getenv("DSN")

	// Open MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	//context
	ctx := h.Context{
		Store: &dao.Store{db},
	}

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
