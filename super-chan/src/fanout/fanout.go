package fanout

import (
    "context"
    "time"
    "fmt"
    "sync"
    "sync/atomic"
)

type Fanout struct {
    src chan string
    dsts []chan string
    parentCtx , currentCtx context.Context
    stopper context.CancelFunc
    isRunning int32
    sendTimeoutMs int32
}

func New(ctx context.Context, src chan string, dsts []chan string, sendTimeoutMs int32) *Fanout {
    return &Fanout {src: src, dsts: dsts, parentCtx: ctx, sendTimeoutMs: sendTimeoutMs}
}

func (fo *Fanout) Run() error {
    if atomic.LoadInt32(&fo.isRunning) == 1 {
        return fmt.Errorf("fanout is already running")
    }
    if err := fo.parentCtx.Err(); err != nil {
        return err
    }
    atomic.StoreInt32(&fo.isRunning, 1)
    
    fo.currentCtx, fo.stopper = context.WithCancel(fo.parentCtx)
    for {
        select {
            case msg, more := <- fo.src:
                var wg sync.WaitGroup
                wg.Add(len(fo.dsts))
                for _, ch := range fo.dsts {
                    go func(ch chan string) {
                        defer wg.Done()
                        err := sendToOneChan(fo.currentCtx, msg, ch, fo.sendTimeoutMs)
                        if err != nil {
                            fmt.Printf("%v\n", err)
                        }
                    }(ch)
                }
                wg.Wait()
                if !more {
                    fo.Stop()
                    return nil
                }
            case <- fo.currentCtx.Done():
                fo.Stop()
                return nil
        }
    }
}

func sendToOneChan(ctx context.Context, msg string, dst chan string, timeoutMs int32) error {
    timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(timeoutMs) * time.Millisecond)
    defer cancel()
    select {
        case dst <- msg:
            return nil
        case <- timeoutCtx.Done():
            return fmt.Errorf("could not send message to %v: timeout %d ms", dst, timeoutMs)
    }
}

func (fo *Fanout) Stop() {
    if fo.stopper != nil {
        fo.stopper()
    }
    atomic.StoreInt32(&fo.isRunning, 0)
}