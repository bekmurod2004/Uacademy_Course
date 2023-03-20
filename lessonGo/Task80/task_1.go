package main

import "fmt"

func Task1() {
	var num int 
	var num2 int 
	var num3 int 
	fmt.Scanln(&num)
	fmt.Scanln(&num2)
	fmt.Scanln(&num3)

	if num < num2 && num < num3 {
		fmt.Println(num)
	}else if num2 < num && num2 < num3 {
		fmt.Println(num2)
	}else if num3 < num && num3 < num2 {
		fmt.Println(num3)
	}else {
		fmt.Println("numbers similar")
	}


}
