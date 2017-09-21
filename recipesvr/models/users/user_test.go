package users

import "testing"

func TestNewUserValidate(t *testing.T) {
	nu := &NewUser{
		Email:        "mail@mail.com",
		Password:     "password",
		PasswordConf: "password",
		FirstName:    "ena",
	}

	// test if fields are there
	err := nu.Validate()
	if err == nil {
		t.Fatalf("no error returned when a field was missing on new user")
	}
	if err != ErrMissingField {
		t.Fatalf("wrong error returned for missing field on new user")
	}

	nu.LastName = "mark"
	err = nu.Validate()
	if err != nil {
		t.Fatalf("validate returned error when fields were passable")
	}

	// test if email is valid
	err = nu.Validate()
	if err != nil {
		t.Fatalf("validate returned error when email was passable")
	}

	nu.Email = "mail.com"
	err = nu.Validate()
	if err == nil {
		t.Fatalf("no error returned when email field was not valid")
	}
	nu.Email = "mail@mail.com"

	// test if password is too short
	err = nu.Validate()
	if err != nil {
		t.Fatalf("validate returned error when passwords is ok length")
	}

	nu.Password = "pass"
	err = nu.Validate()
	if err == nil {
		t.Fatalf("no error when password is too short")
	}
	if err != ErrPasswordTooShort {
		t.Fatalf("wrong error returned when password is too short")
	}

	// test if passwords match
	nu.PasswordConf = "pass"
	err = nu.Validate()
	if err != ErrPasswordTooShort {
		t.Fatalf("wrong error returned when password is too short")
	}

	nu.Password = "password"
	nu.PasswordConf = "password"

	err = nu.Validate()
	if err != nil {
		t.Fatalf("validate returned error when passwords match")
	}

	nu.PasswordConf = "sdfkjhsd"
	err = nu.Validate()
	if err == nil {
		t.Fatalf("no error returned when passwords do not match")
	}
	if err != ErrPasswordsDontMatch {
		t.Fatalf("wrong error returned for passwords not matching")
	}
}
