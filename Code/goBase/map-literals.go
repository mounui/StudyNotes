package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": Vertex{
        40.42343, -75.43543,
    },
    "Google": Vertex{
        37.45323, -122.43542,
    },
}

func main() {
    fmt.Println(m)
}
