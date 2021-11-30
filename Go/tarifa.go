package main

import "fmt"

func main() {
        var x, n, total int
        fmt.Scanln(&x)
        fmt.Scanln(&n)

        for i := 0; i < n; i++ {
                var p int
                fmt.Scanln(&p)
                total += p
        }

        fmt.Println((n+1)*x - total)
}