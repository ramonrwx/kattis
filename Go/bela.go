package main

import "fmt"

// dominant and not dominant
var dNd = map[byte][]int{
	'A': []int{11, 11},
	'K': []int{4, 4},
	'Q': []int{3, 3},
	'J': []int{20, 2},
	'T': []int{10, 10},
	'9': []int{14, 0},
	'8': []int{0, 0},
	'7': []int{0, 0},
}

func dinput(d byte) int {
	var input string
	fmt.Scanln(&input)

	if input[1] == d {
		return dNd[input[0]][0]
	} else {
		return dNd[input[0]][1]
	}
}

func main() {
	var c, sum int
	var d byte
	fmt.Scanf("%d %c\n", &c, &d)

	c <<= 2
	for i := 0; i < c; i++ {
		sum += dinput(d)
	}
	fmt.Println(sum)
}
