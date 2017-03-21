package models

import (
	"gopkg.in/mgo.v2/bson"
)

//gravatarBasePhotoURL is the base URL for Gravatar profile photos
const gravatarBasePhotoURL = "https://www.gravatar.com/avatar/"

//User represents a user account in the database
type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Email     string        `json:"email"`
	PassHash  []byte        `json:"-" bson:"passHash"` //stored in mongo, but never encoded to clients
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	PhotoURL  string        `json:"photoURL"`
}

//Credentials represents user sign-in credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//NewUser represents a new user signing up for an account
type NewUser struct {
	Credentials         //embeds the fields from Credentials in this struct
	PasswordConf string `json:"passwordConf"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

//Validate validates the new user
func (nu *NewUser) Validate() error {
	//ensure Email field is a valid Email
	//HINT: use mail.ParseAddress()
	//https://golang.org/pkg/net/mail/#ParseAddress

	//ensure Password is at least 6 chars

	//ensure Password and PasswordConf match

	//ensure LastName has non-zero length

	//if you made here, it's valid, so return nil
	return nil
}

//ToUser converts the NewUser to a User
func (nu *NewUser) ToUser() (*User, error) {
	//build the Gravatar photo URL by creating an MD5
	//hash of the new user's email address, converting
	//that to a hex string, and appending it to their base URL:
	//https://www.gravatar.com/avatar/ + hex-encoded md5 has of email

	//construct a new User setting the various fields
	//use bson.NewObjectId() to generate a new ID for the user

	//call the User's SetPassword() method to set the password,
	//which will hash the plaintext password

	//return the User and nil
	return nil, nil
}

//SetPassword hashes the password and stores it in the PassHash field
func (u *User) SetPassword(password string) error {
	//hash the plaintext password using an adaptive
	//crytographic hashing algorithm like bcrypt
	//https://godoc.org/golang.org/x/crypto/bcrypt

	//set the User's PassHash field to the resulting hash

	return nil
}

//Authenticate compares the plaintext password against the stored hash
//and returns an error if they don't match, or nil if they do
func (u *User) Authenticate(password string) error {
	//compare the plaintext password with the PassHash field
	//using the same hashing algorithm you used in SetPassword

	return nil
}