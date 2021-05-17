package main

import (
    "fmt"
)

func main() {
    var workArray [10]uint8
    var temp uint8
    for i := 0; i < len(workArray); i++ {
        fmt.Scan(&temp)
        workArray[i] = temp
    }
    var swapA, swapB int
    for swapCounter := 0; swapCounter < 3; swapCounter ++ {
        fmt.Scan(&swapA)
        fmt.Scan(&swapB)
        if swapA == swapB {
            continue
        }
        temp = workArray[swapA]
        workArray[swapA] = workArray[swapB]
        workArray[swapB] = temp
    }
    for _, elem := range workArray {
        fmt.Printf("%d ", elem)
    }
}