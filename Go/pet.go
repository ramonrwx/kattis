package main

import (
        "fmt"
)

func main() {
        var a, b, c, d, tmp, total, winner int

        for i := 1; i <= 5; i++ {
                fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

                tmp = a + b + c + d
                if tmp > total {
                        winner = i
                        total = tmp
                }
        }

        fmt.Println(winner, total)
}