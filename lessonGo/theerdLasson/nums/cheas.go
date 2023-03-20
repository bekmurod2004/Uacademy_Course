package nums

import (
	"fmt"
)


func NumBum(num int) {
	if num < 16 && num > 0{
		ans := 0

	for i := 1; i < 10; i++ {
		
		ans = num * i
		fmt.Printf("%v x %v = %v\n",num ,i, ans)
		
		
	}
		
	}else {
		for i := 1; i <= 9; i++ {
			ans := 0

	for j := 1; j < 10; j++ {
		
		ans = i * j
		fmt.Printf("%v x %v = %v\t",i ,j, ans)
		
		
	}
	fmt.Println()
			
		}
	}
	







	

	
	

}
