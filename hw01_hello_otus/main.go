package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	reverseText := reverse.String("Hello, OTUS!")
	fmt.Println(reverseText)
}
