package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

var stdOutWriter io.Writer

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(input []string, env Environment) (returnCode int) {
	if err := setEnv(env); err != nil {
		log.Fatal(err)
	}
	if stdOutWriter == nil {
		SetOutWriter(os.Stdout)
	}
	commands := createCommands(input)
	for _, cmd := range commands {
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, cmd := range commands {
		err := cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
	}
	return commands[len(commands)-1].ProcessState.ExitCode()
}

func SetOutWriter(w io.Writer) {
	stdOutWriter = w
}

func setEnv(env Environment) error {
	for name, eItem := range env {
		if eItem.NeedRemove {
			if err := os.Unsetenv(name); err != nil {
				return err
			}
			continue
		}
		if err := os.Setenv(name, eItem.Value); err != nil {
			return err
		}
	}
	return nil
}

func createCommands(input []string) []*exec.Cmd {
	var in io.ReadCloser
	in = os.Stdin
	out := []*exec.Cmd{}
	var cmd *exec.Cmd
	for _, v := range input {
		v = strings.Trim(v, " ")
		if v == "|" {
			if cmd != nil {
				out = append(out, cmd)
				cmd = nil
			}
			continue
		}
		if cmd == nil {
			cmd = exec.Command(v)
			cmd.Stdin = in
			in, _ = cmd.StdoutPipe()
			continue
		}
		if cmd != nil {
			cmd.Args = append(cmd.Args, v)
		}
	}
	if cmd != nil {
		out = append(out, cmd)
	}
	if len(out) != 0 {
		out[len(out)-1].Stdout = stdOutWriter
	}
	return out
}
