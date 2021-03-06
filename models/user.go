package models

import (
	"os"
	"user_auth/Credentials"
	"user_auth/helpers"
	userModel "user_auth/models/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IUserModel interface {
	Signup(cred Credentials.SignUpCredentials) error
	GetUserByEmail(email string) (user userModel.User, err error)
}

type UserModel struct {
	Db      *mgo.Database
	session *mgo.Session
}

// Signup handles registering a user
func (u *UserModel) Signup(cred Credentials.SignUpCredentials) error {
	// Connect to the user collection
	collection := u.Db.C("user")
	// Assign result to error object while saving user
	err := collection.Insert(bson.M{
		"name":     cred.Name,
		"email":    cred.Email,
		"password": helpers.GeneratePasswordHash([]byte(cred.Password)),
	})

	return err
}

// GetUserByEmail handles fetching user by email
func (u *UserModel) GetUserByEmail(email string) (user userModel.User, err error) {
	// Connect to the user collection
	collection := u.Db.Session.DB(os.Getenv("DBNAME")).C("user")
	// Assign result to error object while saving user
	err = collection.Find(bson.M{"email": email}).One(&user)
	return user, err
}
