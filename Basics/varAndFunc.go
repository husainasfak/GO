package main

// Format
import "fmt"

// Package level varible
var myName string

func main(){
	fmt.Println("hello world")

	// Declaring a varibale -- function level variable

	var whatToSay string

	whatToSay = "Hey, how are you"

	fmt.Println(whatToSay)

	var i int = 20;

	fmt.Println(i)


	somthing := saySomthing() // colon set what is returning from function and its type

	println(somthing)


	thing1,thing2 := sayTwoThings()

	println(thing1,thing2)
}


func saySomthing() string{
	return "I am saying somthing"
}

// return multiple things

func sayTwoThings() (string,string){
	return "thing 1","thing 2"
}