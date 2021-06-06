package main

import (
	"fmt"
    "strings"
)

type Battery uint

func (b *Battery) String() string {
    result := "[" + strings.Repeat(" ", 10 - int(*b)) + strings.Repeat("X", int(*b)) + "]"
    return result
}

func main() {
	var input string
    fmt.Scan(&input)
    var charge uint
    for _, r := range input {
        if r == '1' {
            charge++
        }
    }
    battery := Battery(charge)
    batteryForTest := &battery
    fmt.Println(batteryForTest)
}