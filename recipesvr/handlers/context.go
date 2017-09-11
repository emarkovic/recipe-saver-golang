package handlers

import (
	"github.com/recipe-saver-golang/recipesvr/models/recipes"
	"github.com/recipe-saver-golang/recipesvr/models/users"
)

// Context represents a struct with helpful information for the handlers.
type Context struct {
	// session id signing key
	// SessionKey string
	// SessionStore sessions.Store
	UserStore   users.Store
	RecipeStore recipes.Store
}
