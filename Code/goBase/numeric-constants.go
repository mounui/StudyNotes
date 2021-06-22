package main

import "fmt"

const (
    // 将1左移 100 位来创建一个非常大的数字
    // 即这个数是二进制1 后面跟着 100 个 0
    Big = 1 << 100
    Small = Big >> 99
)

func needInt(x int) int { return x * 10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}
