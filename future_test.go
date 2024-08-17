package ego

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFutureOkay(t *testing.T) {
	f := Go(func() int {
		time.Sleep(time.Millisecond * 10)
		return 4242
	})
	require.Equal(t, 4242, f.Wait())
}

func TestFutureErr(t *testing.T) {
	v, err := Try(func() string {
		return Go(func() string {
			panic("this failed")
		}).Wait()
	})
	require.Equal(t, "", v)
	require.Error(t, err)
	require.Contains(t, err.Error(), "this failed")
}
