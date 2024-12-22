package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var terminalZero, lineBreak []byte = []byte{0}, []byte{10}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	dirItems, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)

	for _, dirItem := range dirItems {
		fileInfo, err := dirItem.Info()
		if err != nil {
			return nil, err
		}

		if fileInfo.IsDir() {
			continue
		}

		if strings.Contains(fileInfo.Name(), "=") {
			continue
		}

		if fileInfo.Size() == 0 {
			env[fileInfo.Name()] = EnvValue{NeedRemove: true}
		}

		value, err := readFile(dir + "/" + fileInfo.Name())
		if err != nil {
			return nil, err
		}

		if len(value) == 0 {
			env[fileInfo.Name()] = EnvValue{NeedRemove: true}
		} else {
			env[fileInfo.Name()] = EnvValue{Value: value}
		}
	}

	return env, nil
}

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer func() {
		errC := file.Close()
		if errC != nil {
			log.Println(errC)
		}
	}()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return "", nil
	}
	str := scanner.Text()

	return trim(str), nil
}

func trim(content string) string {
	strBytes := []byte(content)
	strBytes = bytes.ReplaceAll(strBytes, terminalZero, lineBreak)
	content = string(strBytes)
	content = strings.TrimRight(content, "\t ")

	return content
}
