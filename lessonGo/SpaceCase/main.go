package main

import "fmt"

// import "fmt"

// import "fmt"

func main() {
	// fmt.Println(t2(5," wo"))
	// Merge()
	// Dupl()
	// UperC()
	var worf string = "testes"
	var sbr = []int{}

	for i := 0; i < len(worf); i++ {
		jum := 0

		for j := 0; j < len(worf); j++ {
			if worf[i] == worf[j] {
				fmt.Print(string(worf[j]))
				jum++
			}
		}
		fmt.Print(jum)
		sbr = append(sbr, jum)
		fmt.Println()
	}

	fmt.Println(sbr)
	for i := 0; i < len(sbr); i++ {
		if sbr[i] == 1 {
			fmt.Println(string(worf[i]))
			return

		}

	}

	// var st string = "Creating kata is fun"
	// kj := 14
	// var br string
	// var str string = "..."

	// for i := 0; i < kj -3; i++ {
	// 	br +=string(st[i])
	// }

	// fmt.Println(br + str)

	// 	var arr = []int{1, 1, 1, 1, 1}

	// 	for i := 0; i < len(arr)-1; i++ {

	// 		for j := i + 1; j < len(arr); j++ {
	// 			arr[i] += arr[j]
	// 		}
	// 		arr[i+1] += arr[i]
	// 	}
	// 	fmt.Println(arr)

	// }

	// 5 9 12
	// 5 1 1 1 1
	// var word string = "the-stealth-warrior"
	// ans := []string{}

	// for i := 0; i < len(word); i++ {
	// 	if string(word[i]) == "-" {
	// 		// fmt.Println(string(word[i +1]))
	// 		i++
	// 		ans = append(ans, strings.ToUpper(string(word[i])))
	// 	} else {
	// 		ans = append(ans, string(word[i]))
	// 	}

	// }

	// fmt.Println(ans)

	// func t2(num int,  s string) string   {
	// 	var satr string

	// 	for i := 0; i < num; i++ {
	// 		satr += s
	//     }

	// 	return satr

	// }
}
