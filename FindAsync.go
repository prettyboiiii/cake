package cake

import (
	"context"
)

/*
Asynchronously find the particular value follow the predicate func in the cake.Slice in a new goroutine
Param:

  - param1: predicate func, its return value indicate what value to find
  - param2: context passed from the main goroutine for managing and propagating cancellation, deadlines, and other request-scoped values

Returns:

	A channel that will send the founded value (`pc` returns `true`)
	if the value does not exist in the cake.Slice or the context is canceled in the process, zero value of T is sent
*/
func (s Slice[T]) FindAsync(pc Predicate[T], ctx context.Context) *FindOp[T] {
	ch := make(chan T, 1)
	return &FindOp[T]{
		srcCh: ch,
		ctx:   ctx,
		execFn: func() {
			defer close(ch)
			for _, v := range s {
				select {
				case <-ctx.Done():
					break
				default:
					if pc(v) {
						ch <- v
						return
					}
				}
			}
			var zeroV T
			ch <- zeroV
		},
	}
}

type FindOp[T any] struct {
	srcCh   chan T
	ctx     context.Context
	execFn  func()
	foundV  T
	founded bool
}

type AsyncOp[T any] interface {
	BgExec() chan T
	Exec() T
}
