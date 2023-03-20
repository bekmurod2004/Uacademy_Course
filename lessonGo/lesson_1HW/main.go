package main

import (
	"fmt"
	"strconv"
)

func main() {

	arr := []int{}

	var a int = 12321

	for a > 0 {
		arr = append(arr,a%10 )
		a = a /10
	}

	fmt.Println(arr)
	

	var word string

	fmt.Scanln(&word)

	pal(word)

	if len(word)%2 == 0 {

		first := string(word[0])
		last := string((word[len(word)-1]))

		f, _ := strconv.Atoi(first)
		l, _ := strconv.Atoi(last)

		if f > l {
			fmt.Println(f)
		}else {
			fmt.Println(l)
		}

		

	} else {
		var str string
		for i := len(word) / 2; i > 0; i-- {
			str = string(word[i])
			break
		}

		first := string(word[0])
		last := string((word[len(word)-1]))


		f, _ := strconv.Atoi(first)
		c, _ := strconv.Atoi(str)
		l, _ := strconv.Atoi(last)

		if f > c && f >l {
			fmt.Println(f)
		}else if c > f && c > l {
			fmt.Println(c)
		}else if l > f && l > c {
			fmt.Println(l)
		}else {
			fmt.Println(word ,"<= have similar")
		}
		
		

	}

}

// var ans string = "hnu"

// for i := 0; i < len(word) -1; i++ {
// 	for j := i + 1; j < len(word); j++ {
// 		if word[i] == word[j] {
// 			ans =  "chiroyli"
//         }
//     }

// 	}

// 	fmt.Println(ans)

// }

// fmt.Println("S =",math.Pow(float64(word), 2))
// fmt.Println( "P =", math.Pow(float64(word), 4))

// pal(word)

// var ans string
// var n2 int64

// powe(int(n1) , int(n2))

// if n1 %2 != 0 && n1 % n1 == 0 {
// 	fmt.Println(true)

// }else{
// 	fmt.Println(false)
// }

// for i := 0; i < int(n1); i++ {
// 	var telme string
// 	fmt.Scanln(&telme)
// 	telme += " "

// 	ans += telme

// }

// fmt.Println(ans)

// fmt.Scanln(&n2)

// fmt.Printf( "type :%T value:%v  ", n1, n1)
// println()
// fmt.Printf("type:%T value:%v  ", n2, n2)

// func powe(a , b int)  {

// 	var i int = 0
// 	var q int = 1

// 	for  i < b {
// 		i++

// 		q = a * a

// 	}

// 	fmt.Println(q)

//
