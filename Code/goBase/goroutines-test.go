package main

import "fmt"

func say(s int) {
    fmt.Println(s)
}

func main() {
    for i := 0; i < 1e4; i++ {
        say(i)
    }
}
