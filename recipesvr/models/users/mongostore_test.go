package users

import "testing"
import "gopkg.in/mgo.v2"
import "reflect"

func TestMongoStore(t *testing.T) {
	sess, err := mgo.Dial("127.0.0.1:27017")
	defer sess.Close()
	if err != nil {
		t.Fatalf("error dialing Mongo: %v", err)
	}

	mgoStore := &MongoStore{
		Session:        sess,
		DatabaseName:   "UsersDB",
		CollectionName: "Users",
	}

	newUser := &NewUser{
		Email:        "test@mail.com",
		Password:     "password",
		PasswordConf: "password",
		FirstName:    "Test",
		LastName:     "TestLast",
	}

	// testing insert
	user, err := mgoStore.Insert(newUser)
	if err != nil {
		t.Fatalf("error inserting user: %v", err)
	}
	if nil == user {
		t.Fatalf("nil user returned from insert")
	}
	if len(user.ID) == 0 {
		t.Fatalf("length of user id is 0")
	}

	// testing get by email
	userByEmail, err := mgoStore.GetByEmail(user.Email)
	if err != nil {
		t.Errorf("error getting by email: %v", err)
	}
	if nil == userByEmail {
		t.Errorf("nil user returned from get by GetByEmail, check function implementation")
	}
	if !reflect.DeepEqual(user, userByEmail) {
		t.Errorf("inserted user and user gotten by inserted user's email are not the same, check GetByEmail implementation")
	}

	userByID, err := mgoStore.GetByID(user.ID)
	if err != nil {
		t.Errorf("error getting by id: %v", err)
	}
	if nil == userByID {
		t.Errorf("nil user returned from GetByID, check function implementation")
	}
	if !reflect.DeepEqual(user, userByID) {
		t.Errorf("inserted user and user gotten by inserted user's ID are not the same, check GetByID implementation")
	}

	mgoStore.Session.DB(mgoStore.DatabaseName).C(mgoStore.CollectionName).RemoveAll(nil)
}
