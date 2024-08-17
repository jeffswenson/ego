package ego

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func frameA(msg string, args ...any) Exception {
	return Errorf(msg, args...)
}

func frameB(msg string, args ...any) Exception {
	return frameA(msg, args...)
}

func frameC(msg string, args ...any) Exception {
	return frameB(msg, args...)
}

func TestFormatException(t *testing.T) {
	e := frameC("this is '%s-%d' the error", "some", 2).(*exception)
	e.callers = e.callers[:4]

	require.Equal(t, expectedStack, fmt.Sprintf("%+v", e))
	require.Equal(t, "this is 'some-2' the error", fmt.Sprintf("%s", e))
}

var expectedStack = `this is 'some-2' the error
0: ego.frameA exception_test.go:11
1: ego.frameB exception_test.go:15
2: ego.frameC exception_test.go:19
3: ego.TestFormatException exception_test.go:23`
