package main

import "fmt"

func main() {
	var cases int
	var sum float64

	fmt.Scanln(&cases)

	for i := 0; i < cases; i++ {
		var period, years float64
		fmt.Scanf("%f %f", &period, &years)

		sum += period * years
	}

	fmt.Printf("%.3f\n", sum)
}
