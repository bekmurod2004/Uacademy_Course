package main

import (
	"strings"
)

func RealValid(name string , lastName string, phone string, email string) (bool , bool , bool) {




	isNameAndLast := len(name) >= 3 && len(lastName) >= 3
	// +998 97 000 00 00
	havePlus := string(phone[0]) == "+" && string(phone[1]) == "9"&& string(phone[2]) == "9"&& string(phone[3]) == "8" && string(phone[4]) == "9"&& len(phone) == 13
	haveDOg := strings.Contains(email , "@")



	// fmt.Println(isNameAndLast)
	// fmt.Println(havePlus)
	// fmt.Println(haveDOg)



	return isNameAndLast, havePlus,haveDOg





	




	
	
}