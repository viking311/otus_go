package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Env = modifyEnv(env)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}

		return 111
	}

	return
}

func modifyEnv(env Environment) []string {
	result := os.Environ()

	for name, value := range env {
		result = removeEnvVariable(result, name)

		if value.NeedRemove {
			continue
		}

		result = append(result, name+"="+value.Value)
	}

	return result
}

func removeEnvVariable(env []string, name string) []string {
	i := 0

	for _, envVar := range env {
		varName := strings.SplitN(envVar, "=", 2)[0]
		if varName != name {
			env[i] = envVar
			i++
		}
	}

	return env[:i]
}
