package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	Slider()

	

	// arr := []int{1,2,3,2,4,1}

	// abc := make(map[int][]int)


	


	// for i := 0; i < len(arr); i++ {
	// 	arm := []int{}
	// 	for j := i+ 1; j < len(arr); j++ {
	// 		if arr[i] == arr[j] {
			
	// 			arm = append(arm, i,j)
	// 			abc[arr[i]] = arm

	// 			// fmt.Println(arr[i], i , j)
				
	// 		}

			
	// 	}


		


		
	// }

	// fmt.Println(abc)

	


	//  m := map[int]int{
    //     0: 1,
    //     1: 2,
    //     2: 1,
    //     3: 4,
    //     4: 2,
    //     5: 7,
    // }
    // mNew := make(map[int]int)

    // for k, v := range m {
	// 	val, has := mNew[v]
		
    //     if has == false {
    //         mNew[v] = k
    //     } else {
    //         fmt.Println(k, m[k], ",", val, m[val])
    //     }
    // }

	


	
	
}


func Slider(){


	


	var inp int

	fmt.Scanln(&inp)

	arr := make([]string, inp)

	lenv := ""

	
	for i := 0; i < len(arr); i++ {
	
		in := bufio.NewReader(os.Stdin)
		arr[i], _ = in.ReadString('\n')



		
		lenv += arr[i]
		
	}

	fmt.Println(arr)
	fmt.Println(len(lenv))





	// arr := []int{1,2,3,4,5,6,7,8,9,10}

	// var (
	// 	n1 = 0
	// 	n2 = 1
	// 	n3 = 2
	// )
 
	// for true {

	// 	n1++
	// 	n2++
	// 	n3++


	// 	if n1 == len(arr) {
	// 		n1 = 0
			
	// 	}

	// 	if n2 == len(arr) {
	// 		n2 = 0
			
	// 	}

	// 	if n3 == len(arr) {
	// 		n3 = 0
			
	// 	}


	// 	fmt.Println(arr[n1], arr[n2],arr[n3])
		
		
	}

	 

	



// 123
// 234
// 345
// 451
// 512
// 123