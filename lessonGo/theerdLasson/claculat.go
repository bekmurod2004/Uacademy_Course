package main

import "fmt"

func Culc(a float64 ,b int, act string)  {
	if act == "+"{
		
		fmt.Println(convertInt(a) + b)
	} else if act == "-" {
		fmt.Println(a - convertFlo(b))
		
	}
	
}

func convertFlo(a int) float64   {
	return float64(a)
}

func convertInt(a float64) int   {
	fmt.Println(int(a))
	return int(a)
}