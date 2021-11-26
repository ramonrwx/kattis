package main

import (
        "fmt"
)

func main() {
        var tBase, tHeight, aTriangle float64
        fmt.Scanf("%f %f", &tBase, &tHeight)

        aTriangle = (tBase * tHeight) / 2
        fmt.Println(aTriangle)
}