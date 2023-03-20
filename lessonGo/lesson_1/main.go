package main

import (
	"fmt"
	"lessonGo/lesson_1/ho"
)

func main() {
	ho.Ism()
	var num1 float64
	var num2 float64
	var opr string

	fmt.Print("enter number 1: ")
	fmt.Scanln(&num1)
	fmt.Print("enter number 2: ")
	fmt.Scanln(&num2)
	fmt.Print("what do you want: (  +   -   *   /  ) => ")
	fmt.Scanln(&opr)

	if opr != "+" && opr != "-" && opr != "*" && opr != "/" {
		fmt.Println("error")
		return
	}

	fmt.Println(calc(num1, num2, opr))

}
