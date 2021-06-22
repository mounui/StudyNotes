package main

import "fmt"

func main() {
    balance := [5]float32{1:2, 3:7.0}
    for k := 0; k < 5; k++ {
        fmt.Printf("balance[%d] = %f\n", k, balance[k])
    }
}
