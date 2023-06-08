package main

import "log"

func main() {
	// For loop, normal
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	// in slice
	animals := []string{"dog","horse","cow","cat","rabbit"}

	for i, animal:=range animals{
		log.Println(i,animal)
	}

	for _, animal:=range animals{
		log.Println(animal)
	}


	// in map
	anime := make(map[string]string)
	anime["luffy"] = "one piece"
	anime["itachi"] = "naruto"

	for animeChar, anime:= range anime{
		log.Println(animeChar,anime)
	}


	// in string
	var str = "Onece upon a time when there was a travelling saleman";
	// in go string is slice of bytes
	for i, s:= range str{
		log.Println(i,":",s);
	}


	// in custom type
	type User struct {
		FirstName string
		LastName string
		age int
		isAdult bool
	}

	var users []User
	users = append(users, User{"Jhon","Doe",16,false})
	users = append(users, User{"Rumi","Roy",31,true})

	for _, u:=range users{
		log.Println(u.FirstName,u.isAdult)
	}

}