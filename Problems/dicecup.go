package main

import "fmt"

func minMax(n, m *int) {
        if *n > *m {
                *n = *n ^ *m
                *m = *m ^ *n
                *n = *n ^ *m
        }
}

func main() {
        var n, m int
        fmt.Scanln(&n, &m)

        minMax(&n, &m)

        for i := n + 1; i < m+2; i++ {
                fmt.Println(i)
        }
}