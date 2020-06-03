package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

const allowMethods = "GET, PUT, POST, PATCH, DELETE"
const allowHeaders = "Content-Type, Authorization"
const exposeHeaders = "Authorization"
const maxAge = "600"

/**
CorsHandler represents a middleware handler for handling CORS requests.
*/
type CorsHandler struct {
	handler http.Handler
}

func (h *CorsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Set headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", allowMethods)
	w.Header().Set("Access-Control-Allow-Headers", allowHeaders)
	w.Header().Set("Access-Control-Expose-Headers", exposeHeaders)
	w.Header().Set("Access-Control-Max-Age", maxAge)

	// Handle preflight request
	if r.Method == "OPTIONS" {
		// Check request method
		reqMethod := r.Header.Get("Access-Control-Request-Method")
		if !strings.Contains(allowMethods, reqMethod) {
			http.Error(w, fmt.Sprintf("Bad CORS preflight request: unsupported method `%s`", reqMethod), http.StatusBadRequest)
			return
		}

		// Check headers
		reqHeaders := strings.Split(r.Header.Get("Access-Control-Request-Headers"), ",")
		for _, header := range reqHeaders {
			if !strings.Contains(strings.ToLower(allowHeaders), strings.ToLower(strings.TrimSpace(header))) {
				http.Error(w, fmt.Sprintf("Bad CORS preflight request: unsupported header `%s`", header), http.StatusBadRequest)
				return
			}
		}

		// Preflight passed, respond
		w.WriteHeader(http.StatusOK)
	} else {
		// Serve actual handler
		h.handler.ServeHTTP(w, r) // future: last tricky part is expectedExuctionOfHandler
	}
}

/**
SetCors wraps a http.Handler with the CORS middleware handler
*/
func SetCors(handlerToWrap http.Handler) *CorsHandler {
	return &CorsHandler{handlerToWrap}
}
