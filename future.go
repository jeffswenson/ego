package ego

// Go is a replacement for the go keyword. It runs the argument function in a
// goroutine and recovers any exception thrown by the function.
func Go[T any](do func() T) *Future[T] {
	f := NewFuture[T]()
	go func() {
		value, err := Try(do)
		f.Complete(value, err)
	}()
	return f
}

// Future is used to track the result of an async task.
type Future[T any] struct {
	done chan struct{}
	t    T
	err  Exception
}

// NewFuture constructs an in-complete future.
func NewFuture[T any]() *Future[T] {
	return &Future[T]{
		done: make(chan struct{}),
	}
}

// Complete must be called when the task is finished. If err is non nil, the
// value is ignored.
func (f *Future[T]) Complete(value T, err Exception) {
	if err != nil {
		f.err = err
	} else {
		f.t = value
	}
	close(f.done)
}

// Wait for the task to complete. Returns a value if the task completed
// successfully or throws an exception if ther was an error.
func (f *Future[T]) Wait() T {
	<-f.done
	AssertNil(f.err)
	return f.t
}

// WaitErr waits for the task to complete. It returns the value and the error.
func (f *Future[T]) WaitErr() (T, Exception) {
	<-f.done
	return f.t, f.err
}

// Complete returns true once the task is complete and false if it is still
// running.
func (f *Future[T]) IsComplete() bool {
	select {
	case <-f.done:
		return true
	default:
		return false
	}
}
