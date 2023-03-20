package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type kars struct{
	Name string  `json:"Name"`
	Price string  `json:"Price"`
}

func ReadJson() {

	dat , err := ioutil.ReadFile("cars.json")

	if err != nil {
		fmt.Println(err)

		return
	}

	var lol []kars

	err = json.Unmarshal(dat ,&lol)

	if err != nil {
		fmt.Println(err)

		return
	}


	for _, v := range v {
		
	}





	
}