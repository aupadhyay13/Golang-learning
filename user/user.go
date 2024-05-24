package user

import "fmt"
import "time"
import "errors"

type User struct {		 // if first character is uppercase then it is globally available
	firstName string
	lastName string
	birthDate string
	createdAt time.Time  // nested struct
} 
func (u User) OutputUserDetails(){   // here u user is called as receiver and it attaches methos to struct
	fmt.Println(u.firstName, u.lastName, u.birthDate)
	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate)// technical way of accessing structs value by pointer

}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User,error){	// constructor kind of function
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil,errors.New("FirstName , LastName and birthDate required")
	}
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}