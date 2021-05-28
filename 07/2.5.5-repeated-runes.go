package main

import (
    "fmt"
    "strings"
    "sort"
)

func main() {
    s := ReadAndTrimSingleNewline()
    existingRunes := make(map[rune]int)
    
    for idx, c := range s {
        if existingRunes[c] > 0 {
            existingRunes[c] = -1
            continue
        }
        _, exists := existingRunes[c]
        if !exists {
            existingRunes[c] = idx + 1
        }
    }
    
    swapped := make(map[int]rune)
    swapKeys := make([]int, 0, len(existingRunes))
    for c, idx := range existingRunes {
        if idx > 0 {
            swapKeys = append(swapKeys, idx)
            swapped[idx] = c
        }
    }
    sort.Ints(swapKeys)
    
    var b strings.Builder
    for _, idx := range swapKeys {
        b.WriteRune(swapped[idx])
    }
    
    fmt.Println(b.String())
}