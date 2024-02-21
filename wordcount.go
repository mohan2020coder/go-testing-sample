package main

import (
    "fmt"
    "strings"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    counts := make(map[string]int)

    for _, word := range words {
        counts[word]++
    }

    return counts
}

func main() {
    s := "hello world hello"
    wc := WordCount(s)
    for word, count := range wc {
        fmt.Printf("%s: %d\n", word, count)
    }
}
