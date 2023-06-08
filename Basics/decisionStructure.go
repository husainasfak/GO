package main

import "log"

func main() {
	var isTrue bool
	myNum := 100
	isTrue = true

	if isTrue &&  myNum > 2{
		log.Println("hey")
	}else if isTrue && myNum > 49{
		log.Println("bye")
	}else{
		log.Println("hey bye")
	}


	switch myNum {
	case 100 :
		log.Println("cat is long and hairy")
	default:
		log.Println("dog")
	}

}