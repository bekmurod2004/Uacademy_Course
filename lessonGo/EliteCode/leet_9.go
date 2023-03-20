package main

import "fmt"

func PalindromNum(x int) bool {
	y := x

	if x <= 0 {
		return false
	}
	temp := 0
	for x > 0 {
		temp *= 10
		temp += x % 10
		x = x / 10

	}
	fmt.Println(temp)
	fmt.Println(x)

	if temp == y {
		return true
	} else if temp != x {
		return false
	}
	return true

}
