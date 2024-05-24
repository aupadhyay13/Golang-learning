package main

import (
	"fmt"
	"time"
)


type user struct {		 // if first character is uppercase then it is globally available
	firstName string
	lastName string
	birthDate string
	createdAt time.Time  // nested struct
} 
func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	var appUser user


	appUser = user {
		firstName: userFirstName,
		lastName: userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}
	outputUserDetails(firstName, lastName, birthdate)
}

func outputUserDetails(firstName,lastName,birthdate string){
	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}