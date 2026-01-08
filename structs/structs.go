package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *user) outputUserDetails() { //special method attached to struct user, can not be called as standalone function anymore
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearUserName() { // required to let this method edit the original Struct... instead of creating a copy and editing
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(text string) string {
	fmt.Println(text)
	var value string
	fmt.Scan(&value)
	return value
}
