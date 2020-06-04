package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/my/repo/ljchen17/final-project/MeetingScheduler/servers/user/model"
	"github.com/my/repo/ljchen17/final-project/MeetingScheduler/servers/user/sessions"
)

/**
UsersHandler handles requests for the "users" resource.
POST requests create new user accounts.
*/
func (ctx *HandlerCtx) UsersHandler(w http.ResponseWriter, r *http.Request) {

	// Check request method
	if r.Method == "POST" {

		// Check content type
		ctype := r.Header.Get("Content-Type")
		if !strings.HasPrefix(ctype, "application/json") {
			http.Error(w, "Unsupported content type "+ctype, http.StatusUnsupportedMediaType)
			return
		}

		// Read response body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Could not read request body", http.StatusBadRequest)

			return
		}
		defer r.Body.Close()

		// Unmarshal response body into new user
		newUser := &model.NewUser{}
		if err := json.Unmarshal(body, newUser); err != nil {
			http.Error(w, "Could not decode request body as user data", http.StatusBadRequest)

			return
		}

		// Validate user
		user, err := newUser.ToUser()
		if err != nil {
			http.Error(w, "Invalid new user", http.StatusBadRequest)
			return
		}

		// Add user to database
		user, err = ctx.UserStore.Insert(user)
		if err != nil {
			status := http.StatusInternalServerError
			if driverErr, ok := err.(*mysql.MySQLError); ok {
				// Handle duplicate email/username error
				if driverErr.Number == 1062 {
					status = http.StatusBadRequest
				}
			}
			http.Error(w, "Error inserting user into database", status)
			return
		}

		// Begin user session
		state := &SessionState{time.Now(), user}
		if _, err := sessions.BeginSession(ctx.SigningKey, ctx.SessionStore, state, w); err != nil {
			http.Error(w, "Error beginning user session", http.StatusInternalServerError)
			return
		}

		// Respond to client
		response, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Error marshalling user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	} else {
		http.Error(w, "Unsupported method "+r.Method, http.StatusMethodNotAllowed)
	}
}

/**
SpecificUserHandler handles requests for a specific user.
GET requests retrieve a user's information.
PATCH requests update a user's information.
*/
func (ctx *HandlerCtx) SpecificUserHandler(w http.ResponseWriter, r *http.Request) {

	// Check request method
	if r.Method != "GET" && r.Method != "PATCH" {
		http.Error(w, "Unsupported method "+r.Method, http.StatusMethodNotAllowed)
		return
	}

	session, _ := sessions.GetSessionID(r, ctx.SigningKey)

	stateRet := SessionState{}
	ctx.SessionStore.Get(session, &stateRet)

	var user *model.User
	var err error
	var userID int64
	urlID := path.Base(r.URL.Path)
	// Get user ID
	if urlID == "me" {
		userID = stateRet.User.UID
	} else {
		userID, err = strconv.ParseInt(urlID, 10, 64)
		if err != nil {
			http.Error(w, "Error parsing user ID", http.StatusBadRequest)
			return
		}
	}

	// Get user from DB
	user, err = ctx.UserStore.GetByID(userID)
	if err == model.ErrUserNotFound {
		http.Error(w, "User with ID "+urlID+" not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error getting user from database", http.StatusInternalServerError)
		return
	}

	if r.Method == "PATCH" {
		// Check if user is authorized
		if urlID != "me" && urlID != strconv.FormatInt(stateRet.User.UID, 10) {
			http.Error(w, "Cannot update other users", http.StatusForbidden)
			return
		}

		// Check content type
		ctype := r.Header.Get("Content-Type")
		if !strings.HasPrefix(ctype, "application/json") {
			http.Error(w, "Unsupported content type "+ctype, http.StatusUnsupportedMediaType)
			return
		}

		// Read response body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Could not read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		updates := &model.Updates{}
		if err := json.Unmarshal(body, updates); err != nil {
			http.Error(w, "Could not decode request body as updates", http.StatusBadRequest)
			return
		}

		// Update user
		user, err = ctx.UserStore.Update(userID, updates)
		if err != nil {
			http.Error(w, "Error updating user", http.StatusInternalServerError)
			return
		}
	}

	// Respond to client
	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshalling user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

/**
SessionsHandler handles requests for the "sessions" resource.
POST requests begin new sessions.
*/
func (ctx *HandlerCtx) SessionsHandler(w http.ResponseWriter, r *http.Request) {

	// Check request method
	if r.Method == "POST" {
		// Check content type
		ctype := r.Header.Get("Content-Type")
		if !strings.HasPrefix(ctype, "application/json") {
			http.Error(w, "Unsupported content type "+ctype, http.StatusUnsupportedMediaType)
			return
		}

		// Read response body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Could not read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		creds := &model.Credentials{}
		if err := json.Unmarshal(body, creds); err != nil {
			http.Error(w, "Could not decode request body as credentials", http.StatusBadRequest)
			return
		}

		// Authenticate user
		user, err := ctx.UserStore.GetByEmail(creds.Email)

		if err == model.ErrUserNotFound {
			time.Sleep(2 * time.Second)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, "Error getting user from database", http.StatusInternalServerError)
			return
		}

		if err := user.Authenticate(creds.Password); err != nil {
			time.Sleep(2 * time.Second)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Begin session
		state := &SessionState{time.Now(), user}

		if _, err = sessions.BeginSession(ctx.SigningKey, ctx.SessionStore, state, w); err != nil {
			http.Error(w, "Error beginning user session", http.StatusInternalServerError)
			return
		}

		// Logs user sign-in attempts
		ipAddress := r.RemoteAddr
		forwardedHeader := r.Header.Get("X-Forwarded-For")
		if forwardedHeader != "" {
			ips := strings.Split(forwardedHeader, ", ")
			ipAddress = ips[0]
		}

		logEntry := &model.LogEntry{
			UserID:     user.UID,
			SignInTime: time.Now(),
			IPAddress:  ipAddress,
		}

		ctx.UserStore.InsertLog(logEntry)

		// Respond to client
		response, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Error marshalling user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	} else {
		http.Error(w, "Unsupported method "+r.Method, http.StatusMethodNotAllowed)
	}
}

/**
SpecificSessionHandler handles requests for a specific authenticated session.
DELETE requests end the current session.
*/
func (ctx *HandlerCtx) SpecificSessionHandler(w http.ResponseWriter, r *http.Request) {

	// Check request method
	if r.Method == "DELETE" {

		// Check request path
		if path.Base(r.URL.Path) != "mine" {
			http.Error(w, "Cannot end other users' sessions", http.StatusForbidden)
			return
		}

		// End session
		_, err := sessions.EndSession(r, ctx.SigningKey, ctx.SessionStore)
		if err != nil {
			http.Error(w, "Error ending session", http.StatusInternalServerError)
			return
		}

		// Write response
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("signed out"))
	} else {
		http.Error(w, "Unsupported method "+r.Method, http.StatusMethodNotAllowed)
	}
}
