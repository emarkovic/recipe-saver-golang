package users

import mgo "gopkg.in/mgo.v2"

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

func (ms *MongoStore) Insert(newUser *NewUser) (*User, error) {
	return nil, nil
}
