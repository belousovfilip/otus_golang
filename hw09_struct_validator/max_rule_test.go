package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxRule(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := MaxRule{Value: "test", Max: "4"}
		passes, err := r.Passes()
		require.NoError(t, err)
		require.True(t, passes)

		r = MaxRule{Value: []int{1, 2, 3}, Max: "4"}
		passes, err = r.Passes()
		require.NoError(t, err)
		require.True(t, passes)
	})
	t.Run("error", func(t *testing.T) {
		r := MaxRule{Value: "test", Max: "3"}
		passes, err := r.Passes()
		require.NoError(t, err)
		require.False(t, passes)

		r = MaxRule{Value: []int{1, 2, 3, 4}, Max: "2"}
		passes, err = r.Passes()
		require.NoError(t, err)
		require.False(t, passes)
	})
}
