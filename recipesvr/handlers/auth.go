package handlers

import "net/http"

// UsersHandler handles creating a new user
// /v1/users -> post
func (ctx *Context) UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "invalid request method", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("you are trying to create a new user"))
}

// SessionsHandlers handles sign ins/outs for existing users
// /v1/sessions -> post -> signs in
// /v1/sessions -> delete -> signs out
func (ctx *Context) SessionsHandlers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	switch r.Method {
	case "POST":
		w.Write([]byte("you are trying to sign in"))
		// create the session
	case "DELETE":
		w.Write([]byte("you are trying to sign out"))
		// delete the session
	}

}
