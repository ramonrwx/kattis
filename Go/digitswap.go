package main

import "fmt"

func main() {
        var input int
        fmt.Scanf("%d", &input)

        fmt.Printf("%v%v", input%10, float64(input/10))
}