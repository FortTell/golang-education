package main

import "fmt"

func main() {
    var d int
    fmt.Scan(&d)
    h := d / 30
    m := (d % 30 ) * 2
    fmt.Println("It is", h, "hours", m, "minutes.")
}