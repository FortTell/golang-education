package main

import (
    "superChan/src/pipe"
    "fmt"
)

func main() {
    cin := make(chan string)
    cout := make(chan string)
    pp := pipe.New(cin, cout, func(s string) string { return s + s })
    go func() {
        pp.Run()
    }()
    
    go func() {
        for i := 0; i < 26; i++ {
            cin <- string('a' + i)
        }
    }()
    for i := 0; i < 13; i++ {
        x := <- cout
        fmt.Println(x)
    }
    pp.Stop()
}