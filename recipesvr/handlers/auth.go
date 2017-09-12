package handlers

import "net/http"

// UsersHandler handles creating a new user
// /v1/users -> post
func (ctx *Context) UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "invalid request method", http.StatusBadRequest)
		return
	}
}

// SessionsHandlers handles sign ins for existing users
// /v1/sessions -> post
func (ctx *Context) SessionsHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "invalid request method", http.StatusBadRequest)
		return
	}
}
