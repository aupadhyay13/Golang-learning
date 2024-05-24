package main

import (
	"fmt"
	"example.com/structs/user"
)



func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	var appUser *user.User

	// appUser = user{} // this will create with null value


	appUser,err := user.New(userFirstName,userLastName,userBirthdate)
	if err != nil{
		fmt.Println(err)
		return
	}
	appUser.OutputUserDetails() // no need to pass argumemnts because it will be passed automatically by go
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}