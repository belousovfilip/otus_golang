package main

import (
	"bytes"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("2 cmd", func(t *testing.T) {
		b := bytes.NewBuffer(make([]byte, 0))
		env, err := ReadDir("./testdata/env")
		require.NoError(t, err)
		SetOutWriter(b)
		rCode := RunCmd([]string{"env", "|", "grep", "BAR"}, env)
		status := syscall.WaitStatus(rCode)
		require.True(t, status.Exited())
		require.Equal(t, "BAR=bar\n", b.String())
	})
	t.Run("3cmd", func(t *testing.T) {
		b := bytes.NewBuffer(make([]byte, 0))
		env, err := ReadDir("./testdata/env")
		require.NoError(t, err)
		SetOutWriter(b)

		rCode := RunCmd([]string{"env", "|", "grep", "BAR", "|", "wc", "-l"}, env)
		status := syscall.WaitStatus(rCode)

		require.Equal(t, "1\n", b.String())
		require.Equal(t, 0, status.ExitStatus())
		require.True(t, status.Exited())
	})
}
