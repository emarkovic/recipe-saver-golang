package handlers

import "net/http"

// RecipesHandler handle api routes for recipe related actions
// /v1/recipes
func (ctx *Context) RecipesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// "saving" -> creating a new recipe
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("You are saving a recipe"))
	case "GET":
		// get all recipe
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("you are getting all recipes"))
	case "PATCH":
		// assign to day of the week
		// favorite
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("you are either assinging a recipe or favoriting it"))
	case "DELETE":
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("you are deleting a recipe"))
	}
}
