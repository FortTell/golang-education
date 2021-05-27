package main

import (
    "fmt"
    "errors"
)

func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    } else {
        return a / b, nil
    }
}

func main() {
    var a, b int
    fmt.Scan(&a)
    fmt.Scan(&b)
    result, err := divide(a, b)
    if err != nil {
        fmt.Println("ошибка")
    } else {
        fmt.Println(result)
    }
}