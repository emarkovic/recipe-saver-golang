package users

import (
	"errors"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// ErrPasswordTooShort is returned when new user password is too short
var ErrPasswordTooShort = errors.New("password is less than 6 characters")

// ErrPasswordsDontMatch is returned when new user password and password confirmation do not match
var ErrPasswordsDontMatch = errors.New("password and password confirmation do not match")

// ErrMissingField is returned when new user did not provide all required fields
var ErrMissingField = errors.New("required field is missing or empty")

//UserID defines the type for user IDs
type UserID string

//User represents a user account in the database
type User struct {
	ID        UserID `json:"id" bson:"_id"`
	Email     string `json:"email"`
	PassHash  []byte `json:"-" bson:"paddHash"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//Credentials represents user sign-in credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

func isBlank(str string) bool {
	str = strings.TrimSpace(str)
	return len(str) == 0
}

//Validate validates the email address and the password and password confirmation.
func (nu *NewUser) Validate() error {
	if isBlank(nu.Email) ||
		isBlank(nu.Password) ||
		isBlank(nu.PasswordConf) ||
		isBlank(nu.FirstName) ||
		isBlank(nu.LastName) {

		return ErrMissingField
	}

	_, err := mail.ParseAddress(nu.Email)
	if err != nil {
		return err
	}

	if len(nu.Password) < 6 {
		return ErrPasswordTooShort
	}

	if nu.Password != nu.PasswordConf {
		return ErrPasswordsDontMatch
	}

	return nil
}

//ToUser converts the NewUser to a User
func (nu *NewUser) ToUser() (*User, error) {
	user := &User{
		Email:     nu.Email,
		FirstName: nu.FirstName,
		LastName:  nu.LastName,
	}

	user.SetPassword(nu.Password)

	return user, nil
}

//SetPassword hashes the password and stores it in the PassHash field
func (u *User) SetPassword(password string) error {
	passhash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return errors.New("error generating password hash")
	}

	u.PassHash = passhash

	return nil
}

//Authenticate compares the plaintext password against the stored hash
//and returns an error if they don't match, or nil if they do
func (u *User) Authenticate(password string) error {
	err := bcrypt.CompareHashAndPassword(u.PassHash, []byte(password))
	if err != nil {
		return errors.New("Invalid credentials")
	}

	return nil
}
