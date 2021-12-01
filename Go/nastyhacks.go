package main

import (
	"fmt"
)

func main() {
	var cases, r, e, c int

	fmt.Scanf("%d", &cases)

	for i := 1; i <= cases; i++ {
		fmt.Scanf("%d %d %d", &r, &e, &c)
		if (e - c) > r {
			fmt.Println("advertise")
		} else if (e - c) < r {
			fmt.Println("do not advertise")
		} else {
			fmt.Println("does not matter")
		}
	}
}
