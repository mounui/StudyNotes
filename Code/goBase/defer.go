package main

import "fmt"
import "time"

func main() {
    defer fmt.Println(time.Now())

    fmt.Println("hello")
}
