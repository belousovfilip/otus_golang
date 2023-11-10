package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInRule(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		r := InRule{Value: "33", list: []string{"33", "44"}}
		passes, err := r.Passes()
		require.NoError(t, err)
		require.True(t, passes)
	})
	t.Run("not found", func(t *testing.T) {
		r := InRule{Value: "11", list: []string{"33", "44"}}
		passes, err := r.Passes()
		require.NoError(t, err)
		require.False(t, passes)
	})
}
