package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    sFields := strings.Fields(s)
    var resMap = map[string]int{}
    for _,v := range sFields{
        resMap[v] ++
    }
    return resMap
}

func main() {
    wc.Test(WordCount)
}
