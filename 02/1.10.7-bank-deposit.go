package main

import (
    "fmt"
)

func main() {
    var money, pctPerYear, desiredMoney int
    fmt.Scan(&money)
    fmt.Scan(&pctPerYear)
    fmt.Scan(&desiredMoney)
    multiplier := 100 + pctPerYear
    for year := 0; ; year++ {
        if money >= desiredMoney {
            fmt.Println(year)
            break
        }
        money = int((money * multiplier) / 100.0)
    }
}

