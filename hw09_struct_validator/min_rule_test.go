package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinRule(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := MinRule{Value: "test", Min: "4"}
		passes, err := r.Passes()
		require.NoError(t, err)
		require.True(t, passes)

		r = MinRule{Value: []int{1, 2, 3, 4, 5}, Min: "4"}
		passes, err = r.Passes()
		require.NoError(t, err)
		require.True(t, passes)
	})
	t.Run("error", func(t *testing.T) {
		r := MinRule{Value: "few", Min: "4"}
		passes, err := r.Passes()
		require.NoError(t, err)
		require.False(t, passes)
		r = MinRule{Value: "Я", Min: "2"}
		passes, err = r.Passes()
		require.NoError(t, err)
		require.False(t, passes)

		r = MinRule{Value: []int{1}, Min: "2"}
		passes, err = r.Passes()
		require.NoError(t, err)
		require.False(t, passes)
	})
}
