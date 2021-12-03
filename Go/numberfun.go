package main

import "fmt"

func main() {
	var cases int
	fmt.Scanln(&cases)

	for i := 0; i < cases; i++ {
		var a, b, c int
		fmt.Scanf("%d %d %d", &a, &b, &c)
		if produce(a, b, c) {
			fmt.Println("Possible")
		} else {
			fmt.Println("Impossible")
		}
	}
}

func produce(a, b, c int) bool {
	ab := a / b
	ba := b / a
	switch {
	case a+b == c:
		return true
	case a-b == c:
		return true
	case b-a == c:
		return true
	case a*b == c:
		return true
	case ab == c && ab*b == a:
		return true
	case ba == c && ba*a == b:
		return true
	default:
		return false
	}
}
