package sessions

import (
	"errors"
	"net/http"
	"strings"
)

const headerAuthorization = "Authorization"
const paramAuthorization = "auth"
const schemeBearer = "Bearer "

// ErrNoSessionID is used when no session ID was found in the Authorization header
var ErrNoSessionID = errors.New("no session ID found in " + headerAuthorization + " header")

// ErrInvalidScheme is used when the authorization scheme is not supported
var ErrInvalidScheme = errors.New("authorization scheme not supported")

/**
BeginSession creates a new SessionID, saves the `sessionState` to the store, adds an
Authorization header to the response with the SessionID, and returns the new SessionID
*/
func BeginSession(signingKey string, store Store, sessionState interface{}, w http.ResponseWriter) (SessionID, error) {

	// Create a new SessionID using signingKey
	sessionID, err := NewSessionID(signingKey)
	if err != nil {
		return InvalidSessionID, err
	}

	// Save the sessionState to the store
	data := store.Save(sessionID, sessionState)
	if data != nil {
		return InvalidSessionID, err
	}

	// Add a header to the ResponseWriter
	w.Header().Add(headerAuthorization, schemeBearer+sessionID.String())

	// Return newly created SessionID
	return sessionID, nil
}

/**
GetSessionID extracts and validates the SessionID from the request headers
*/
func GetSessionID(r *http.Request, signingKey string) (SessionID, error) {

	// Get the value of the Authorization header;
	// if no Authorization header is present, get the "auth" query string parameter
	auth := r.Header.Get(headerAuthorization)
	if auth == "" {
		auth = r.URL.Query().Get(paramAuthorization)
	}
	if auth == "" {
		return InvalidSessionID, ErrNoSessionID
	}

	// Validate the auth obtained
	if !strings.HasPrefix(auth, schemeBearer) {
		return InvalidSessionID, ErrInvalidScheme
	}
	auth = strings.TrimPrefix(auth, schemeBearer)

	// Validate the SessionID; if it's not valide, return the validation error
	sessionID, err := ValidateID(auth, signingKey)
	if err != nil {
		return InvalidSessionID, ErrNoSessionID
	}

	return sessionID, nil
}

/**
GetState extracts the SessionID from the request,
gets the associated state from the provided store into
the `sessionState` parameter, and returns the SessionID
*/
func GetState(r *http.Request, signingKey string, store Store, sessionState interface{}) (SessionID, error) {

	// Get the SessionID from the request
	sessionID, err := GetSessionID(r, signingKey)
	if err != nil {
		return InvalidSessionID, err
	}

	// Get the data associated with that SessionID from the store
	data := store.Get(sessionID, sessionState)
	if data != nil {
		return InvalidSessionID, data
	}

	return sessionID, nil
}

/**
EndSession extracts the SessionID from the request,
and deletes the associated data in the provided store, returning
the extracted SessionID.
*/
func EndSession(r *http.Request, signingKey string, store Store) (SessionID, error) {

	// Get the SessionID from the request
	sessionID, err := GetSessionID(r, signingKey)
	if err != nil {
		return InvalidSessionID, err
	}

	// Delete the data associated with it in the store
	err = store.Delete(sessionID)
	if err != nil {
		return InvalidSessionID, err
	}

	return sessionID, nil
}
