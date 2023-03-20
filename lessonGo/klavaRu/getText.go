package main

import "fmt"

func ReadText() {
	about := []string{"b - б   v - в   g - г   d - д" , "j - ж   z - з   k - к   l - л", "m - м   n - н   p - п   r - р" ,"s - c   t - т   f - ф   h - х", "y - й   a - a   e - e   i - и ", "o - o   u - y   ц - ?   щ - $","ё - :   ы - /   э - )   ю - #","я - %   ь - |   ъ - &   [ - ч  " ,"] - ш"}


	for i := 0; i < len(about); i++ {
		if i == 4 {
			fmt.Println(about[i])
			fmt.Println()
		}else {
			fmt.Println(about[i])
		}
		
		
	}

	fmt.Println()






var theWord string


fmt.Print("Write Russian text on English letter: ")

action := true

all := ""

for action{
	
	
	fmt.Scan(&theWord)

	all += theWord
	all += " "
	



	if theWord == "..." {
		action = false
	}


}





fmt.Println("-----------------------------------------------")


fmt.Print("Russian text on Russian letter[ ")
Translation(all)
fmt.Println("]")

fmt.Println("-----------------------------------------------")

	
}