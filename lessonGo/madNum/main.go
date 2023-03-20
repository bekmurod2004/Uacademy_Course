package main

import (
	"fmt"
	"strconv"
)

func main() {

	action := true

	var num string
	
	var ans int
	
	for action {
		fmt.Scanln(&num)
		bum, err  := strconv.Atoi(num)
		ans += bum
		if err != nil || num == "stop" {
			fmt.Println(ans)
			action = false
		}

	}
	// var nome string

	// fmt.Scanln(&nome)

	// // 123456 9  - 12
	// a := 0
	// b := 0

	// for i := 0; i < len(nome); i++ {
	// 	if i %2 == 0 {
	// 		bum ,_ := strconv.Atoi(string(nome[i]))
	// 		a += bum

	// 	}else {
	// 		bum ,_ := strconv.Atoi(string(nome[i]))
	// 		b += bum

	// 	}

	// }

	// fmt.Println(a -b)

	// var number int
	// fmt.Scan(&number)

	// ans := 0

	// for number > 0 {
	// 	ans = ans * 10
	// 	ans += number % 10

	// 	number = number / 10

	// }

	// fmt.Println(ans)

	// var sum int	 = 0
	// for i := 1; i < number; i++ {
	// 	if number %i == 0 {
	// 		sum = sum + i

	// 	}
	// }

	// if number == sum {
	// 	fmt.Println("unique num")
	// }else {
	// 	fmt.Println("no unique")
	// }

}
