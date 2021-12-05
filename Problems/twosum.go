package main

import (
        "fmt"
)

func sum(a, b int) int {
        return a + b
}

func main() {
        var a, b int
        fmt.Scanf("%d %d", &a, &b)

        fmt.Printf("%d", sum(a, b))
}