package main

import "fmt"
import "strconv"

type IPAddr [4]byte

func (i IPAddr) String() string {
    str := ""
    for _, v := range i{
        if str != "" {
            str += "."
        }
        fmt.Printf(strconv.Itoa(int(v)))
        str += strconv.Itoa(int(v))
    }
    return str
}

func main() {
    hosts := map[string]IPAddr{
        "loopback": {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
