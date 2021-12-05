package main

import "fmt"

func main() {
        var n, lastTime, total int
        var running bool
        fmt.Scanln(&n)

        for i := 0; i < n; i++ {
                var currentTime int
                fmt.Scanln(&currentTime)

                if running == false {
                        running = true
                        lastTime = currentTime
                } else {
                        running = false
                        total += (currentTime - lastTime)
                }
        }

        if running {
                fmt.Println("still running")
        } else {
                fmt.Println(total)
        }
}