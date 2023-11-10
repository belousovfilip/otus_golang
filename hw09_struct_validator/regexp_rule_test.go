package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegexpRule(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := RegexpRule{Value: "123", Regexp: `\d+`}
		passes, err := r.Passes()
		require.NoError(t, err)
		require.True(t, passes)
	})
}
