package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

const (
	NullByte      = 0
	LineBreakByte = 10
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var ErrInvalidFileName = errors.New("invalid file name")

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	envMap := make(Environment)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		fInfo, _ := f.Info()
		if fInfo.IsDir() {
			continue
		}
		envName := f.Name()
		if strings.Contains(envName, "=") {
			return nil, ErrInvalidFileName
		}
		file, _ := os.Open(dir + "/" + envName)
		line, _, _ := bufio.NewReader(file).ReadLine()
		for i, b := range line {
			if b == NullByte || b == LineBreakByte {
				line[i] = []byte("\n")[0]
			}
		}
		envValue := strings.TrimRight(string(line), " ")
		envName = strings.TrimSpace(envName)
		envMap[envName] = EnvValue{
			Value:      envValue,
			NeedRemove: len(envValue) == 0,
		}
	}
	return envMap, nil
}
