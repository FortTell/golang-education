package main

import (
    "superChan/src/pipe"
    "superChan/src/fanout"
    "fmt"
    "context"
    "sync"
    "time"
)

func pipeMain() {
    cin := make(chan string)
    cout := make(chan string)
    pp := pipe.New(context.TODO(), cin, cout, func(s string) string { return s + s })
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

func fanoutMain() {
    cin := make(chan string)
    dsts := [2]chan string{}
    for i := 0; i < 2; i++ {
        dsts[i] = make(chan string)
    }
    fo := fanout.New(context.TODO(), cin, dsts[:], 500)
    go func() {
        fo.Run()
    }()
    var wg sync.WaitGroup
    wg.Add(1)
    time.Sleep(1 * time.Second)
    wg.Done()
    go func() {
        for i := 0; i < 26; i++ {
            wg.Add(2)
            cin <- string('a' + i)
        }
    }()
    go func() {
        for {
            msg := <- dsts[0]
            fmt.Printf("channel 0 got %s\n", msg)
            wg.Done()
        }
    }()
    go func() {
        for {
            msg := <- dsts[1]
            fmt.Printf("channel 1 got %s\n", msg)
            wg.Done()
            time.Sleep(2 * time.Second)
        }
    }()
    wg.Wait()
    fo.Stop()
}

func main() {
    fanoutMain()
}