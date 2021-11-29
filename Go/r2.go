package main

import "fmt"

func main() {
        var r1, s int
        fmt.Scanf("%d %d", &r1, &s)

        // (r1 + r2) / 2 = s => 2s = r1 + r2 => r2 = 2s - r1
        fmt.Println((s * 2) - r1) // (s << 1) - r1
}