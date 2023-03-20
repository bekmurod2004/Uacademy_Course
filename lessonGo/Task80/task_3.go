package main

import "fmt"

func Task3() {
	var num int 
	var num2 int 

	fmt.Print("first num: ")
	fmt.Scanln(&num)
	fmt.Print("second num: ")
	fmt.Scanln(&num2)

	if num < num2  {
		x := num
		he := (num + num2 )/2
		num = he

		hu := (x + num2 )*2
		num2 = hu
		
	}else  if  num2 < num {
		y := num2
		hu := (num2 + num )/2
		num2 = hu

		he := (y * num)*2
		num = he
		
	}else {
		fmt.Printf("first num: %v",num)
		fmt.Printf("second num: %v",num2)
		
	}

	fmt.Printf("first num: %v ",num)
	fmt.Println()
	fmt.Printf("second num: %v",num2)
	fmt.Println()
	
}