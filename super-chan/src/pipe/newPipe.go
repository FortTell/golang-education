package pipe

import (
    "context"
)

type NewPipe struct {
    src, dst chan string
    transform func(string) string
    ctx context.Context
    stopper context.CancelFunc
}

func New(src, dst chan string, transform func(string) string) *NewPipe {
    pp := NewPipe {src: src, dst: dst, transform: transform, ctx: context.TODO() }
    return &pp
}

func (pp *NewPipe) Run() {
    pp.ctx, pp.stopper = context.WithCancel(pp.ctx)
    for {
        select {
            case v, more := <- pp.src:
                pp.dst <- pp.transform(v)
                if !more {
                    return
                }
            case <- pp.ctx.Done():
                return
        }
    }
}

func (pp *NewPipe) Stop() {
    pp.stopper()
}