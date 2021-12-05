// in python is more easy
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		cases int
		lines []string
	)

	s := bufio.NewScanner(os.Stdin)
	fmt.Scanln(&cases)

	for i := 0; i < cases; i++ {
		s.Scan()
		lines = append(lines, s.Text())
	}

	testCases := make(map[int][]int)
	for j := 0; j < len(lines); j++ {
		var listStrips []int
		powerStrips := strings.Split(lines[j], " ")
		for _, strips := range powerStrips {
			strip, _ := strconv.Atoi(strips)
			listStrips = append(listStrips, strip)
		}
		testCases[j] = listStrips
	}

	for _, tc := range testCases {
		var outlets int
		for _, t := range tc[1:] {
			outlets += t - 1
		}
		fmt.Println(outlets + 1)
	}
}
