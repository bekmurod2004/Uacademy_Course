package main

import (
	"fmt"
	"strings"
)

func Translation(arr string) {

	russ := []string{"б", "в", "г", "д", "ж", "з", "к", "л", "м", "н", "п", "р", "с", "т", "ф", "х", "й", "a", "e", "и", "o", "у", "ц", "щ", "ё", "ы", "э", "ю", "я", "ь", "ъ", "ч", "ш"}

	summRu := []string{}

	for i := 0; i < len(arr)-4; i++ {
		
		switch string(arr[i]) {
		case " ":
			summRu = append(summRu, " ")
			break
		case "b":
			summRu = append(summRu, russ[0])
			break
		case "v":
			summRu = append(summRu, russ[1])
			break
		case "g":
			summRu = append(summRu, russ[2])
			break
		case "d":
			summRu = append(summRu, russ[3])
			break
		case "j":
			summRu = append(summRu, russ[4])
			break
		case "z":
			summRu = append(summRu, russ[5])
			break
		case "k":
			summRu = append(summRu, russ[6])
			break
		case "l":
			summRu = append(summRu, russ[7])
			break
		case "m":
			summRu = append(summRu, russ[8])
			break
		case "n":
			summRu = append(summRu, russ[9])
			break
		case "p":
			summRu = append(summRu, russ[10])
			break
		case "r":
			summRu = append(summRu, russ[11])
			break
		case "s":
			summRu = append(summRu, russ[12])
			break
		case "t":
			summRu = append(summRu, russ[13])
			break
		case "f":
			summRu = append(summRu, russ[14])
			break
		case "h":
			summRu = append(summRu, russ[15])
			break
		case "y":
			summRu = append(summRu, russ[16])
			break
		case "a":
			summRu = append(summRu, russ[17])
			break
		case "e":
			summRu = append(summRu, russ[18])
			break
		case "i":
			summRu = append(summRu, russ[19])
			break
		case "o":
			summRu = append(summRu, russ[20])
			break
		case "u":
			summRu = append(summRu, russ[21])
			break
		case "?":
			summRu = append(summRu, russ[22])
			break
		case "$":
			summRu = append(summRu, russ[23])
			break
		case ":":
			summRu = append(summRu, russ[24])
			break
		case "/":
			summRu = append(summRu, russ[25])
			break
		case ")":
			summRu = append(summRu, russ[26])
			break
		case "#":
			summRu = append(summRu, russ[27])
			break
		case "%":
			summRu = append(summRu, russ[28])
			break
		case "|":
			summRu = append(summRu, russ[29])
			break
		case "&":
			summRu = append(summRu, russ[30])
			break
		case "[":
			summRu = append(summRu, russ[31])
			break
		case "]":
			summRu = append(summRu, russ[32])
			break
		default:
			summRu = append(summRu, string(arr[i]))
			break

		}

	}

	javob := strings.Join(summRu, "")
	fmt.Print(javob)

}
