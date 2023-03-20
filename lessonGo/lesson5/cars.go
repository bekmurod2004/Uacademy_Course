package main

import (
	"fmt"
	"reflect"
)

type car struct{
	name string
	model string
	maxSpeed int
}

func About() {

	Buggati :=car{
		name: "Buggati",
		model: "g500",
		maxSpeed: 800,
	}

	help := reflect.ValueOf(Buggati)

	fmt.Println(help)
	help1 := []string{"name: ", "model: ", "maxSpeed: "}



	for i := 0; i < help.NumField(); i++ {
		fmt.Print(help1[i])
		fmt.Println(help.Field(i))
		
	}
}