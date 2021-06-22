package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": {40.3242,-74.43214},
    "Google": {37.43234, -122.03483},
}

func main() {
    fmt.Println(m)
}
