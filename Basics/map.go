package main

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "simson"

	println(myMap["dog"])

	myMap2 := make(map[string]int)

	myMap2["fisrt"] = 1

	me := make(map[string]User)
	user := User{
		FirstName: "Ben",
		LastName:  "boom",
	}

	me["me"] = user

	println(me["me"].FirstName)

	// Maps are immutable
	// its like an Object -> key value pair

	mp := make(map[string]interface{}) // take any type of value

	mp["hey"] = 1
	mp["bye"] = "sad"

}
