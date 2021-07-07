package pipe

import (
    "errors"
    "context"
    "sync/atomic"
)

type Pipe struct {
    src, dst chan string
    transform func(string) string
    parentCtx, currentCtx context.Context
    stopper context.CancelFunc
    isRunning int32
}

func New(ctx context.Context, src, dst chan string, transform func(string) string) *Pipe {
    pp := Pipe {src: src, dst: dst, transform: transform, parentCtx: ctx }
    return &pp
}

func (pp *Pipe) Run() error {
    if atomic.LoadInt32(&pp.isRunning) == 1 {
        return errors.New("pipe is already running")
    }
    if err := pp.parentCtx.Err(); err != nil {
        return err
    }
    atomic.StoreInt32(&pp.isRunning, 1)
    
    pp.currentCtx, pp.stopper = context.WithCancel(pp.parentCtx)
    for {
        select {
            case v, more := <- pp.src:
                pp.dst <- pp.transform(v)
                if !more {
                    return nil
                }
            case <- pp.currentCtx.Done():
                return nil
        }
    }
    return nil
}

func (pp *Pipe) Stop() {
    pp.stopper()
    atomic.StoreInt32(&pp.isRunning, 0)
}