package recipes

import "gopkg.in/mgo.v2"

// MongoStore is a struct representing mongo
type MongoStore struct {
	Session        *mgo.Session
	DatabaseName   string
	CollectionName string
}

// NewMongoStore gets a new mongo store.
func NewMongoStore(sessionID *mgo.Session) *MongoStore {
	return &MongoStore{
		Session:        sessionID,
		DatabaseName:   "RecipesDB",
		CollectionName: "Recipes",
	}
}

// GetRecipes gets all recipes a given user is allowed to see
func (ms *MongoStore) GetRecipes() ([]*Recipe, error) {
	return nil, nil
}

// InsertRecipe inserts a new recipe into the db
func (ms *MongoStore) InsertRecipe(newRecipe *NewRecipe) (*Recipe, error) {
	return nil, nil
}

// FavoriteRecipe favorites a recipe
func (ms *MongoStore) FavoriteRecipe(updates *RecipeUpdates, id RecipeID) (*Recipe, error) {
	return nil, nil
}

// ScheduleRecipe saves which day the user applied to the recipe
func (ms *MongoStore) ScheduleRecipe(updates *RecipeUpdates, id RecipeID, day Day) (*Recipe, error) {
	return nil, nil
}

// DeleteRecipe removes a recipe the user has previously saved
func (ms *MongoStore) DeleteRecipe(id RecipeID) error {
	return nil
}
