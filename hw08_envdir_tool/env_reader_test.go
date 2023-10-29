package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("success env", func(t *testing.T) {
		envMap := Environment{
			"BAR":            EnvValue{Value: "bar", NeedRemove: false},
			"EMPTY":          EnvValue{Value: "", NeedRemove: true},
			"FOO":            EnvValue{Value: "   foo\nwith new line", NeedRemove: false},
			"HELLO":          EnvValue{Value: "\"hello\"", NeedRemove: false},
			"UNSET":          EnvValue{Value: "", NeedRemove: true},
			"TAB_AT_THE_END": EnvValue{Value: "TAB", NeedRemove: false},
		}
		list, err := ReadDir("./testdata/env")
		require.NoError(t, err)
		for key, item := range envMap {
			v := list[key]
			require.Equal(t, item, v)
		}
	})
	t.Run("fail env", func(t *testing.T) {
		_, err := ReadDir("./testdata/fail_env")
		require.ErrorIs(t, ErrInvalidFileName, err)
	})
}
