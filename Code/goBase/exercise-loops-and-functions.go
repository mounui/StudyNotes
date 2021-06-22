package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := float64(1)
    for math.Abs(z * z - x) > 1e-10 {
        z -= (z * z - x) / (2 * z)
    }
    return z
}

func Test(x float64) float64 {
    z := 0.1
    for i := 0; i < 10; i++ {
        z -= (z * z - x) / (2 * z)
        fmt.Println(z)
    }
    return z
}

func main() {
    Test(2)
}
