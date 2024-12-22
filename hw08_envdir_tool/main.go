package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatalf("too few arguments")
	}

	env, err := ReadDir(args[1])
	if err != nil {
		log.Fatal(err)
	}

	returnCode := RunCmd(args[2:], env)

	os.Exit(returnCode)
}
