package recipes

import (
	"time"

	"github.com/recipe-saver-golang/recipesvr/models/users"
)

// Day type represents an int
type Day int

// These are the days of the week
const (
	MONDAY    Day = 1
	TUESDAY   Day = 2
	WEDNESDAY Day = 3
	THURSDAY  Day = 4
	FRIDAY    Day = 5
	SATURDAY  Day = 6
	SUNDAY    Day = 7
)

// RecipeID represents the type for recipe IDs
type RecipeID string

// Recipe represents a recipe in the database
type Recipe struct {
	ID        RecipeID     `json:"id" bson:"_id"`
	URL       string       `json:"url"`
	Deleted   bool         `json:"deleted"`
	Day       Day          `json:"day, omitempty" bson:",omitempty"`
	CreatedAt time.Time    `json:"createdAt"`
	CreatorID users.UserID `json:"creatorId"`
}

// NewRecipe represents a user saving a new recipe
type NewRecipe struct {
	RecipeURL string `json:"url"`
}

// RecipeUpdates represents updates a user can make to a saved recipe
type RecipeUpdates struct {
	Favorite bool `json:"favorite"`
	Day      Day  `json:"day"`
	Deleted  bool `json:"deleted"`
}

// ToRecipe converts a NewRecipe to a Recipe
func (nr *NewRecipe) ToRecipe(id users.UserID) *Recipe {
	return &Recipe{
		URL:       nr.RecipeURL,
		Deleted:   false,
		CreatedAt: time.Now(),
		CreatorID: id,
	}
}
