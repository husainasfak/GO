package main

func main() {
	var myString string = "green"

	println(myString)

	changeusingPointer(&myString)

	println(myString)
}

func changeusingPointer(s *string) {
	println("string address", s)
	newValue := "red"
	*s = newValue
}