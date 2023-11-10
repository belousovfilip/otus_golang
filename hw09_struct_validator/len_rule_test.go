package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLenule(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := LenRule{Value: "test", Length: "4"}
		passes, err := r.Passes()
		require.NoError(t, err)
		require.True(t, passes)

		r = LenRule{Value: []int{1, 2, 3, 4}, Length: "4"}
		passes, err = r.Passes()
		require.NoError(t, err)
		require.True(t, passes)
	})
	t.Run("error", func(t *testing.T) {
		r := LenRule{Value: "test", Length: "3"}
		passes, err := r.Passes()
		require.NoError(t, err)
		require.False(t, passes)

		r = LenRule{Value: []int{1, 2, 3, 4}, Length: "2"}
		passes, err = r.Passes()
		require.NoError(t, err)
		require.False(t, passes)
	})
}
