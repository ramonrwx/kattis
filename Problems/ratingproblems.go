package main

import "fmt"

func main() {
        var (
                n, sum float64
                k      int
        )

        fmt.Scanf("%f %d", &n, &k)
        for i := 0; i < k; i++ {
                var r float64
                fmt.Scanln(&r)
                sum += r
        }

        lower := ((sum - (n-float64(k))*3) / n)
        high := ((sum + (n-float64(k))*3) / n)
        fmt.Println(lower, high)
}