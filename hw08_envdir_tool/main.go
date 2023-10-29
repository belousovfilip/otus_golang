package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	envPath := args[0]
	cmd := args[1:]
	env, err := ReadDir(envPath)
	if err != nil {
		log.Fatal(err)
	}
	code := RunCmd(cmd, env)
	log.Printf("Command finished with code: %v", code)
}
