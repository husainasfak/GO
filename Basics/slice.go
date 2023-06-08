package main

import "log"

func main() {
	var mySlice []string

	mySlice = append(mySlice, "trevor")
	mySlice = append(mySlice, "jhon")
	mySlice = append(mySlice, "j")

	log.Println(mySlice)

	// short hand
	numbers := []int{1,2,3,4,5,6}

	log.Println(numbers[0:2])
}