package main

import (
	"fmt"
	"time"
	"errors"
)


type user struct {		 // if first character is uppercase then it is globally available
	firstName string
	lastName string
	birthDate string
	createdAt time.Time  // nested struct
} 
func (u user) outputUserDetails(){   // here u user is called as receiver and it attaches methos to struct
	fmt.Println(u.firstName, u.lastName, u.birthDate)
	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate)// technical way of accessing structs value by pointer

}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthDate string) (*user,error){	// constructor kind of function
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil,errors.New("FirstName , LastName and birthDate required")
	}
	return &user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	var appUser *user

	// appUser = user{} // this will create with null value


	appUser,err := newUser(userFirstName,userLastName,userBirthdate)
	if err != nil{
		fmt.Println(err)
		return
	}
	appUser.outputUserDetails() // no need to pass argumemnts because it will be passed automatically by go
	appUser.clearUserName()
	appUser.outputUserDetails()

}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}