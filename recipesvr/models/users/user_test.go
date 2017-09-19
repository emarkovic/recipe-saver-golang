package users

import "testing"

// type NewUser struct {
// 	Email        string `json:"email"`
// 	Password     string `json:"password"`
// 	PasswordConf string `json:"passwordConf"`
// 	FirstName    string `json:"firstName"`
// 	LastName     string `json:"lastName"`
// }
func TestNewUserValidate(t *testing.T) {
	nu1 := &NewUser{
		Email:        "mail@mail.com",
		Password:     "password",
		PasswordConf: "password",
		FirstName:    "ena",
	}

	// test if fields are there
	// test if email is valid
	// test if passwords match

}
