package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("counting")
    fmt.Println(time.Now())

    var h int = 1
    for i := 0; i < 1e6; i++ {
        h++
    }
    // fmt.Println(h)

    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")
}
