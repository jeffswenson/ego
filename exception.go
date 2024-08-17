package ego

import (
	"fmt"
	"path"
	"runtime"
)

type Exception interface {
	error
	Callers() []uintptr
}

func Errorf(fmtStr string, a ...any) Exception {
	var callers [32]uintptr
	frames := runtime.Callers(2, callers[:])
	return &exception{
		error:   fmt.Errorf(fmtStr, a...),
		callers: callers[:frames],
	}
}

type exception struct {
	error
	callers []uintptr
}

func addStack(skip int, err error) Exception {
	if exception, ok := err.(Exception); ok {
		return exception
	}

	var callers [32]uintptr
	frames := runtime.Callers(2+skip, callers[:])
	return &exception{
		error:   err,
		callers: callers[:frames],
	}
}

func (e *exception) Format(f fmt.State, c rune) {
	switch c {
	case 'v':
		if f.Flag('+') {
			// Print the error message and the stack trace
			fmt.Fprintf(f, "%v\n", e.error)
			frames := runtime.CallersFrames(e.callers)
			for index := 0; ; index++ {
				frame, more := frames.Next()
				if more {
					fmt.Fprintf(f, "%d: %s %s:%d\n", index, path.Base(frame.Function), path.Base(frame.File), frame.Line)
				} else {
					fmt.Fprintf(f, "%d: %s %s:%d", index, path.Base(frame.Function), path.Base(frame.File), frame.Line)
					return
				}
			}
		}
		fallthrough
	case 's':
		fmt.Fprintf(f, "%s", e.Error())
	case 'q':
		fmt.Fprintf(f, "%q", e.Error())
	}
}

func (e *exception) Unwrap() error {
	return e.error
}

func (e *exception) Callers() []uintptr {
	return e.callers
}
