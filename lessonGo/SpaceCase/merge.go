package main

import "fmt"

func Merge() {
	var hu int = 0
	var num int
	a := "abaab"
	b := "aabab"

	for i := 0; i < len(b); i++ {

		if a[i] == b[hu] {

			num = i

		}
	}

	fmt.Println(num)

	var lo string

	for i := 0; i < num-1; i++ {
		lo += string(a[i])

	}

	fmt.Println(lo + b)

}
