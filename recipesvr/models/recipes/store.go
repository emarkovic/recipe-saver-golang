package recipes

import "errors"

// ErrRecipeNotFound is returned when the requested recipe is not found in the store.
var ErrRecipeNotFound = errors.New("recipe not found")

// Store represents an abstract store for model.Recipe objects
// This interface is used by the HTTP handlers to insert new recipes,
// get previously inserted recipes and modify and delete inserted recipes.
// This interface can be implemented for any persistant database
type Store interface {
	// GetRecipes gets all recipes a given user is allowed to see
	// TODO: pass userId
	GetRecipes() ([]*Recipe, error)

	// InsertRecipe inserts a new recipe into the db
	// TODO: pass userId
	InsertRecipe(newRecipe *NewRecipe) (*Recipe, error)

	// FavoriteRecipe favorites a recipe
	FavoriteRecipe(updates *RecipeUpdates, id RecipeID) (*Recipe, error)

	// ScheduleRecipe saves which day the user applied to the recipe
	ScheduleRecipe(updates *RecipeUpdates, id RecipeID, day Day) (*Recipe, error)

	// DeleteRecipe removes a recipe the user has previously saved
	DeleteRecipe(id RecipeID) error
}
