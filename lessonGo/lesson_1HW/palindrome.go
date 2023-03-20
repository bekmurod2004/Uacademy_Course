package main

import "fmt"

func pal(words string) {


	var soz string

	for i := len(words) - 1; i > -1; i-- {
		soz += string(words[i])
		// slik = append(slik, string(words[i]))

	}

	if soz == words {
		fmt.Println(words, "<= palindrom")
	} else {
		fmt.Println(words, "<= no palindrom")

	}

}
