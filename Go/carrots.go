package main

import (
        "fmt"
)

func main() {
        var contests, carrots int
        var desc string

        fmt.Scanf("%d %d", &contests, &carrots)
        for i := 1; i <= contests; i++ {
                fmt.Scanf("%s", &desc)
        }
        fmt.Printf("%d", carrots)
}