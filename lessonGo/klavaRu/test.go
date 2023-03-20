package main

import "fmt"
type Point struct {
	name string
	paswords int
}
func Test() {
	arm := []Point{}

	for i := 0; i <= 10; i++ {
		usr := Point{
		 	name: "test",
			paswords: i,	
		}
		arm = append(arm, usr)	
	}
	fmt.Println(arm)	
}