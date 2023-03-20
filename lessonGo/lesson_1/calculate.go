package main

func calc(n1 float64, n2 float64, op string) float64 {

	var ans float64
	if op == "+" {
		ans = n1 + n2
	} else if op == "-" {
		ans = n1 - n2
	} else if op == "*" {
		ans = n1 * n2
	} else if op == "/" {
		ans = n1 / n2
	}

	return ans
}
