package main

import (
        "fmt"
)

func main() {
        var (
                cost  float64
                lawns int
                total float64 = 0
        )

        fmt.Scanln(&cost)
        fmt.Scanln(&lawns)

        for i := 0; i < lawns; i++ {
                var x, y float64
                fmt.Scanf("%f %f", &x, &y)
                total += x * y
        }

        fmt.Println(total * cost)
}