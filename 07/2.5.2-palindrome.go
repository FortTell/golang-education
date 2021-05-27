package main

import (
    "fmt"
)

func main() {
    text := ReadAndTrimSingleNewline()
    runes := []rune(text)
    
    for strlen, i := len(runes), 0; i < strlen / 2; i++ {
        if runes[i] != runes[strlen - i - 1] {
            fmt.Println("Нет")
            return
        }
    }
    fmt.Println("Палиндром")
}