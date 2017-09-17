package users

import mgo "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"

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

// Insert inserts a new NewUser into the store
// and returns a User with a newly-assigned ID
func (ms *MongoStore) Insert(newUser *NewUser) (*User, error) {
	user, err := newUser.ToUser()
	if err != nil {
		return nil, err
	}

	user.ID = UserID(bson.NewObjectId().Hex())
	err = ms.Session.DB(ms.DatabaseName).C(ms.CollectionName).Insert(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetByEmail returns a user with the provided email address
func (ms *MongoStore) GetByEmail(email string) (*User, error) {
	user := &User{}

	err := ms.Session.DB(ms.DatabaseName).C(ms.CollectionName).Find(bson.M{"email": email}).One(user)
	if err == mgo.ErrNotFound {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

// GetByID returns a user with the provided user id
func (ms *MongoStore) GetByID(userID UserID) (*User, error) {
	user := &User{}

	err := ms.Session.DB(ms.DatabaseName).C(ms.CollectionName).FindId(string(userID)).One(user)
	if err == mgo.ErrNotFound {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
