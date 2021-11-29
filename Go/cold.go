package main

import "fmt"

func main() {
        var n, t, ng_t int
        fmt.Scanln(&n)

        for i := 0; i < n; i++ {
                fmt.Scan(&t)
                if t < 0 {
                        ng_t++
                }
        }
        fmt.Print(ng_t)
}