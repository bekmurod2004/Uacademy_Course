package main

import "fmt"

func main() {

	

	var num int 
    fmt.Scan(&num)
	ans := 0
    

	for num > 0 {
		ans = ans * 10
		ans += num %10
		num = num / 10
	}
	fmt.Println(ans)

	
}