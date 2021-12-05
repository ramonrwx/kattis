package main

import "fmt"

func main() {
	var n, tmp, counter int
	var sum float64

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		if tmp < 0 {
			continue
		}

		counter++
		sum += float64(tmp)
	}

	fmt.Printf("%.10f", sum/float64(counter))
}
