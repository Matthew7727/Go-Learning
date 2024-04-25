package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *user) clearUsername() {
	u.firstName = "" 
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) user {
	return user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthdate,
		createdAt: time.Now().UTC(),
	}
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser = newUser(firstName, lastName, birthdate)

	appUser.outputUserDetails()
	appUser.clearUsername()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

