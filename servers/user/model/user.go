package model

import (
	"fmt"
	"net/mail"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//bcryptCost is the default bcrypt cost to use when hashing passwords
var bcryptCost = 13

//User represents a user account in the database
type User struct {
	UID       int64  `json:"uid"`
	Email     string `json:"email"`
	PassHash  []byte `json:"-"` //never JSON encoded/decoded
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type LogEntry struct {
	UserID     int64
	SignInTime time.Time
	IPAddress  string
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
	UserName     string `json:"userName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

//Updates represents allowed updates to a user profile
type Updates struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

/**
Validate validates the new user and returns an error if
any of the validation rules fail, or nil if its valid
*/
func (nu *NewUser) Validate() error {

	// Check on email field
	_, err := mail.ParseAddress(nu.Email)
	if err != nil {
		return fmt.Errorf("It is not a valid email address")
	}

	// Check on the length of password
	if len(nu.Password) < 6 {
		return fmt.Errorf("Password must be at least 6 characters")
	}

	// Check on the repeated password
	if nu.Password != nu.PasswordConf {
		return fmt.Errorf("Different from the original password")
	}

	// Check on the username
	if len(nu.UserName) == 0 || strings.Contains(nu.UserName, " ") {
		return fmt.Errorf("It is not a valid username")
	}

	return nil
}

/**
ToUser converts the NewUser to a User, setting the
PassHash fields appropriately
*/
func (nu *NewUser) ToUser() (*User, error) {

	// Validate the new user
	err := nu.Validate()
	if err != nil {
		return nil, err
	}

	/*
	   Trim leading and trailing whitespace from an email address
	   Force all characters to lower-case
	   md5 hash the final string
	*/

	// Create a new instance for *User
	user := User{
		UID:       0,
		Email:     nu.Email,
		UserName:  nu.UserName,
		FirstName: nu.FirstName,
		LastName:  nu.LastName,
	}

	err = user.SetPassword(nu.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

/**
FullName returns the user's full name, in the form:
"<FirstName> <LastName>"
If either first or last name is an empty string, no
space is put between the names. If both are missing,
this returns an empty string
*/
func (u *User) FullName() string {

	// Check if either the FirstName or the LastName is empty string
	// Return different form of string accordingly

	if u.FirstName == "" && u.LastName == "" {
		return ""
	} else if u.FirstName == "" {
		return u.LastName
	} else if u.LastName == "" {
		return u.FirstName
	}

	return u.FirstName + " " + u.LastName
}

/**
SetPassword hashes the password and stores it in the PassHash field
*/
func (u *User) SetPassword(password string) error {

	// Generate a new hash of the new user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return err
	}

	// Set the hashed password in the PassHash field of *User
	u.PassHash = hashedPassword

	return nil
}

/**
Authenticate compares the plaintext password against the stored hash
and returns an error if they don't match, or nil if they do
*/
func (u *User) Authenticate(password string) error {

	// Compare the hashed password with its possible plaintext equivalent
	err := bcrypt.CompareHashAndPassword(u.PassHash, []byte(password))

	if err != nil {
		return err
	}

	return nil
}

/**
ApplyUpdates applies the updates to the user. An error
is returned if the updates are invalid
*/
func (u *User) ApplyUpdates(updates *Updates) error {

	// Check the error of the updates
	if updates == nil {
		return fmt.Errorf("The update is not valid")
	}

	if updates.FirstName == "" && updates.LastName == "" {

		return fmt.Errorf("The update is not valid")
	}

	// Apply updates
	u.FirstName = updates.FirstName
	u.LastName = updates.LastName

	return nil
}
