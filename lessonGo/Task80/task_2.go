package main

import "fmt"

func Task2() {

	var num int 
	fmt.Scanln(&num)

	if num %4 == 0 && num %100 != 0 || 	num %400 == 0{
		fmt.Println("366")
		return
	}else {
		fmt.Println("365")
		return
	}

	
	
}