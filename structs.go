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

	// appUser = user{} // this will create with null value


	appUser = user {
		firstName: userFirstName,
		lastName: userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}
	outputUserDetails(&appUser)
}

func outputUserDetails(u *user){
	fmt.Println(u.firstName, u.lastName, u.birthDate)
	fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate)// technical way of accessing structs value by pointer

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}