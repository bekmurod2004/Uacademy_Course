package mapsa

import "fmt"

func Maplash() {
	liblab := make(map[string]string)

	liblab["ArmandoLucas Correa"] = "The Daughter,"
	liblab["HirihikoAraki"] = "Khimyo Naboken Jjba,"
	liblab["AlexandrPushkin"] = "Fisher and fish,"

	fmt.Println(liblab)

	fmt.Print("What do you want 'delet' or 'find' = ")
	var a string
	fmt.Scanln(&a)
	var b string
	b = "find"

	var name string
	var del string

	if a == "delete" {
		fmt.Print("tell me what delete:")
		fmt.Scanln(&del)

		for i, _ := range liblab {

			if i == del {
				delete(liblab,i)
				

			}

		}
		fmt.Println(liblab)
		return





	} else if a == b {
		fmt.Print("tell me name:")
		fmt.Scanln(&name)

		for i, finder := range liblab {

			
			

			if i == name {
				fmt.Println(finder)
				

			}

		}
		fmt.Println(liblab)
		return

	}

}
