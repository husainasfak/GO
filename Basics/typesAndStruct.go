package main

import (
	"time"
)

var s = "seven"

type User struct{
	FirstName string
	LastName string
	age int
	bday time.Time
}

// Assing a func to struct User
func (m *User) PrintFirstName() string{
	return m.FirstName
}

func main() {

	// log.Println(s)

	user := User {
		FirstName: "bon",
		
	}
	var myVar User
	myVar.LastName = "bom";
	println("user last name is ",user.LastName)
	println("user last name is ",myVar.LastName)
	println("user first name is ",user.PrintFirstName())
}

func saySomthing(s string) (string,string){
	return s, "world"
}