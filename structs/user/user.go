package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u *User) OutputUserDetails() { //special method attached to struct user, can not be called as standalone function anymore
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() { // *User required to let this method edit the original Struct... instead of creating a copy and editing
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "ADMIN",
			birthDate: "-----",
			createdAt: time.Now(),
		},
	}
}

// e.g. creation / constructor function
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}
