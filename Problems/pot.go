package main

import (
        "fmt"
        "math"
)

func main() {
        var n int
        var total, p float64

        fmt.Scanln(&n)

        for i := 1; i <= n; i++ {
                fmt.Scanln(&p)
                mPow := math.Mod(p, 10)
                p = math.Floor(p / 10)
                pPow := math.Pow(p, mPow)
                total += math.Round(pPow)
        }

        fmt.Println(int(total))
}