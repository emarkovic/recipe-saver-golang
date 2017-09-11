package users

import "errors"

//ErrUserNotFound is returned when the requested user is not found in the store
var ErrUserNotFound = errors.New("user not found")

//ErrEmailAlreadyExists is returned when the email on insert already exists in the store
var ErrEmailAlreadyExists = errors.New("email is already taken")

//Store represents an abstract store for model.User objects.
//This interface is used by the HTTP handlers to insert new users,
//get users, and update users. This interface can be implemented
//for any persistent database you want (e.g., MongoDB, PostgreSQL, etc.)
type Store interface {
	//Insert inserts a new NewUser into the store
	//and returns a User with a newly-assigned ID
	Insert(newUser *NewUser) (*User, error)
}
