package handler

import (
	"MeetingScheduler/servers/group/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// common used http strings
var contentType = "Content-Type"
var typeJSON = "application/json"

func isContentTypeJSON(w http.ResponseWriter, r *http.Request) bool {
	ctype := r.Header.Get(contentType)
	if ctype != typeJSON {
		errMsg := fmt.Sprintf("Invalid content type [%s], expect [%s]\n", ctype, typeJSON)
		log.Printf(errMsg)
		http.Error(w, errMsg, http.StatusUnsupportedMediaType)
		return false
	}
	return true
}

// TBD: header
func getCurrentUser(w http.ResponseWriter, r *http.Request) int64 {
	xuser := r.Header.Get("X-user")

	if len(xuser) == 0 {
		errMsg := fmt.Sprintf("User not registered\n")
		log.Printf(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return -1
	}

	id, err := strconv.ParseInt(xuser, 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to convert [%s] to id: %v\n", xuser, err)
		log.Printf(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return -1
	}

	return id
}

func getIDfromURL(w http.ResponseWriter, r *http.Request, urlID string) int64 {
	//urlID := path.Base(path.Dir(r.URL.Path))
	id, err := strconv.ParseInt(urlID, 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to convert [%s] to id: %v\n", urlID, err)
		log.Printf(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return -1
	}
	return id
}

func isGroupCreator(group *model.Group, uid int64, w http.ResponseWriter) bool {
	if group.CreatorID != uid {
		errMsg := fmt.Sprintf("The creator of current group is [%d]. User [%d] is not authorized.\n", group.CreatorID, uid)
		log.Printf(errMsg)
		http.Error(w, errMsg, http.StatusUnauthorized)
		return false
	}
	return true
}

func isMeetingCreator(meeting *model.Meeting, uid int64, w http.ResponseWriter) bool {
	if meeting.CreatorID != uid {
		errMsg := fmt.Sprintf("The creator of current group is [%d]. User [%d] is not authorized.\n", meeting.CreatorID, uid)
		log.Printf(errMsg)
		http.Error(w, errMsg, http.StatusUnauthorized)
		return false
	}
	return true
}

func getRequestBody(w http.ResponseWriter, r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errMsg := fmt.Sprintf("Error when trying to read request body: %v\n", err)
		log.Printf(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return nil
	}
	// log.Printf("Get body: %s\n", string(body))
	defer r.Body.Close()
	return body
}

func unmarshalBody(w http.ResponseWriter, body []byte, target interface{}) bool {
	if err := json.Unmarshal(body, target); err != nil {
		errMsg := fmt.Sprintf("Error when trying to unmarshal request body: %v\n", err)
		log.Printf(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return false
	}
	return true
}

func marshalRep(w http.ResponseWriter, obj interface{}) []byte {
	response, err := json.Marshal(obj)
	if err != nil {
		errMsg := fmt.Sprintf("Error when trying to encode JSON: %v\n", err)
		log.Printf(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return nil
	}
	return response
}

func respondWithHeader(w http.ResponseWriter, ctType string, response []byte, code int) {

	w.Header().Set(contentType, ctType)
	w.WriteHeader(code)
	w.Write(response)
	// log.Printf("Respond success")
}

func dbErrorHandle(w http.ResponseWriter, errorPart string, err error) bool {
	if err != nil {
		errMsg := fmt.Sprintf("Error when [%s] in database: %v\n", errorPart, err)
		log.Printf(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return false
	}
	return true
}
