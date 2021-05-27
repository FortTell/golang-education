package main

import (
    "fmt"
)

func runeInRange(check, start, end rune) bool {
    return check >= start && check <= end
}

func main() {
    text := ReadAndTrimSingleNewline()
    if len(text) < 5 {
        fmt.Println("Wrong password")
        return
    }
    
    for _, c := range text {
        if runeInRange(c, '0', '9') || runeInRange(c, 'a', 'z') || runeInRange(c, 'A', 'Z') {
            continue
        }
        fmt.Println("Wrong password")
        return
    }
    fmt.Println("Ok")
}