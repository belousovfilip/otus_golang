package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopyFile(t *testing.T) {
	const input = "./testdata/input.txt"
	tempDir, _ := os.MkdirTemp("", "hw07.")
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println(err)
		}
	}(tempDir)

	t.Run("offset0_limit0", func(t *testing.T) {
		rPath := input
		wPath := tempDir + "/offset0_limit0.txt"
		require.NoError(t, Copy(rPath, wPath, 0, 0))
		rB, _ := os.ReadFile("./testdata/out_offset0_limit0.txt")
		wB, _ := os.ReadFile(wPath)
		require.Equal(t, string(rB), string(wB))
	})

	t.Run("offset0_limit10", func(t *testing.T) {
		rPath := input
		wPath := tempDir + "/out_offset0_limit10.txt"
		require.NoError(t, Copy(rPath, wPath, 0, 10))
		rB, _ := os.ReadFile("./testdata/out_offset0_limit10.txt")
		wB, _ := os.ReadFile(wPath)
		require.Equal(t, string(rB), string(wB))
	})

	t.Run("offset0_limit1000", func(t *testing.T) {
		rPath := input
		wPath := tempDir + "/out_offset0_limit1000.txt"
		require.NoError(t, Copy(rPath, wPath, 0, 1000))
		rB, _ := os.ReadFile("./testdata/out_offset0_limit1000.txt")
		wB, _ := os.ReadFile(wPath)
		require.Equal(t, string(rB), string(wB))
	})

	t.Run("offset0_limit10000", func(t *testing.T) {
		rPath := input
		wPath := tempDir + "/out_offset0_limit10000.txt"
		require.NoError(t, Copy(rPath, wPath, 0, 10000))
		rB, _ := os.ReadFile("./testdata/out_offset0_limit10000.txt")
		wB, _ := os.ReadFile(wPath)
		require.Equal(t, string(rB), string(wB))
	})

	t.Run("offset100_limit1000", func(t *testing.T) {
		rPath := input
		wPath := tempDir + "/out_offset100_limit1000.txt"
		require.NoError(t, Copy(rPath, wPath, 100, 1000))
		rB, _ := os.ReadFile("./testdata/out_offset100_limit1000.txt")
		wB, _ := os.ReadFile(wPath)
		require.Equal(t, string(rB), string(wB))
	})

	t.Run("offset6000_limit1000", func(t *testing.T) {
		rPath := input
		wPath := tempDir + "/out_offset6000_limit1000.txt"
		require.NoError(t, Copy(rPath, wPath, 6000, 1000))
		rB, _ := os.ReadFile("./testdata/out_offset6000_limit1000.txt")
		wB, _ := os.ReadFile(wPath)
		require.Equal(t, string(rB), string(wB))
	})

	t.Run("large_limit", func(t *testing.T) {
		rPath := "./testdata/out_offset0_limit10.txt"
		wPath := tempDir + "/large_limit.txt"
		require.NoError(t, Copy(rPath, wPath, 0, 1000))
		rB, _ := os.ReadFile("./testdata/out_offset0_limit10.txt")
		wB, _ := os.ReadFile(wPath)
		require.Equal(t, string(rB), string(wB))
	})

	t.Run("unknown length", func(t *testing.T) {
		rPath := "/dev/urandom"
		wPath := tempDir + "/unknown_length.txt"
		require.Error(t, ErrUnsupportedFile, Copy(rPath, wPath, 0, 0))
	})
}
