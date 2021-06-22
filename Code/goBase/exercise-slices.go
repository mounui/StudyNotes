package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    out := make([][]uint8, dy)
    for i := 0; i < dy; i++ {
        out[i] = make([]uint8, dx)
        for j := 0; j < dx; j++ {
            out[i][j] = uint8(i*j)
        }
    }
    return out
}

func main() {
    pic.Show(Pic)
}
