package ego

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTryCatchErr(t *testing.T) {
	thrownErr := Errorf("thrown error")
	value, err := Try(func() bool {
		panic(thrownErr)
	})
	t.Log(fmt.Sprintf("%+v", thrownErr))
	require.Equal(t, err, thrownErr)
	require.False(t, value)
}

func TestTryCatchString(t *testing.T) {
	value, err := Try(func() int {
		panic("this is an error")
	})
	require.Containsf(t, err.Error(), "this is an error", "unexpected error message")
	require.Equal(t, value, 0)
}

func TestTryReturnNil(t *testing.T) {
	throwOnly := func() {
		panic("some error")
	}
	value, err := Try(func() any {
		throwOnly()
		return nil
	})
	require.Error(t, err)
	require.Nil(t, value)
}

func TestAssertNilOnNil(t *testing.T) {
	value, err := Try(func() int {
		AssertNil(nil)
		return 42
	})
	require.Nil(t, err)
	require.Equal(t, value, 42)
}

func TestAssertNilOnError(t *testing.T) {
	thrownErr := Errorf("thrown error")
	value, err := Try(func() int {
		AssertNil(thrownErr)
		return 42
	})
	require.Equal(t, err, thrownErr)
	require.Equal(t, value, 0)
}

func TestUnwrapOk(t *testing.T) {
	classicDoer := func(a int) (int, error) {
		return a, nil
	}
	require.Equal(t, Unwrap(classicDoer(1337)), 1337)
}

func TestUnwrapErr(t *testing.T) {
	returnedErr := errors.New("thrown error")
	classicDoer := func(a int) (int, error) {
		return a, returnedErr
	}

	value, err := Try(func() int {
		return Unwrap(classicDoer(1337))
	})
	require.ErrorIs(t, err, returnedErr)
	require.Equal(t, value, 0)
}
