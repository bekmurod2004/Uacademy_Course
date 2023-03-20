package main

import "fmt"

func Dupl() {
	 word := "Helol world!!!H"
	
	

	 for i := 0; i < len(word); i++ {
		
	 var num = -1
		
		 art :=  string(word[i]) 

		 for j := 0; j <len(word); j++ {
			if string(word[j]) == art{
			num++
			if num > 0 {
				fmt.Println(j)
				return
			}	
		 }	
	 }
// bekmurod_
	}

// fmt.Println(num)
}