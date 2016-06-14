package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    res := make(map[string]int)
    tokens := strings.Fields(s)

    for _, token := range tokens {
        if count, exist := res[token]; exist {
            res[token] = count + 1
        } else {
            res[token] = 1
        }
    }

    return res
}

func main() {
    wc.Test(WordCount)
}
