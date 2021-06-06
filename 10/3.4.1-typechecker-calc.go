package main

import (
	"fmt"
    "errors"
)

func readTask() (value1, value2, operation interface{}) {
    return true, 5.6, "/" 
}

func tryCastFloat64(value interface{}) (float64, error) {
    if valueCasted, ok := value.(float64); ok {
        return valueCasted, nil
    } else {
        return 0.0, errors.New(fmt.Sprintf("value=%v: %T", value, value))
    }
}

func main() {
	value1Raw, value2Raw, operationRaw := readTask()
    value1, err := tryCastFloat64(value1Raw)
    if err != nil {
        fmt.Println(err)
        return
    }
    value2, err := tryCastFloat64(value2Raw)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    if operation, ok := operationRaw.(string); ok {
        var result float64
        switch operation {
            case "+": result = value1 + value2
            case "-": result = value1 - value2
            case "*": result = value1 * value2
            case "/": result = value1 / value2
            default:
                fmt.Println("неизвестная операция")
                return
        }
        fmt.Printf("%.4f", result)
    } else {
        fmt.Println("неизвестная операция")
        return
    }
}