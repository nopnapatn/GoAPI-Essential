package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nopnapatn/go-module/hello"
)

func main() {
	// print
	fmt.Print("Hello World\n")

	// use package
	id := uuid.New()
	fmt.Printf("UUID: %s\n", id)

	// use another function
	hello.SayHelloFromPublic()
	hello.SayHelloFromPrivate()

	// but can't use function from internal folder
}
