package recipes

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
	//what goes in here exactly?
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
func (nr *NewRecipe) ToRecipe() *Recipe {
	return &Recipe{}
}
