package main

import (
	"fmt"
	"strings"
	"unicode"
)

func UperC() {
	var word string = "codE"// cODE coDE
	up := 0
	lo := 0
	for _, r := range word {
		if unicode.IsUpper(r) {
			up++
		}
		if unicode.IsLower(r) {
			lo++
		}
	}
	fmt.Println(up)
	fmt.Println(lo)

	if up > lo {
		fmt.Println(strings.ToUpper(word))
	}else  if  lo > up  || lo == up {
		fmt.Println(strings.ToLower(word))
	}
	
}