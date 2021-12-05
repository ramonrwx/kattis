package main

import "fmt"

func main() {
	var (
		n    int
		b, p float64
	)

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var x, y, min, max float64
		fmt.Scanf("%f %f", &b, &p)

		x = 60.0 / p
		y = x * b

		min, max = y-x, y+x

		fmt.Printf("%.4f %.4f %.4f\n", min, y, max)
	}
}
