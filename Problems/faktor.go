package main

import "fmt"

func main() {
        var publish, impact int
        fmt.Scanf("%d %d", &publish, &impact)
        fmt.Println(publish*(impact-1) + 1)
}