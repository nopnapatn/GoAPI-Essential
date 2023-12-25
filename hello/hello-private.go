package hello

import (
	"fmt"
)

func generateHello() {
	fmt.Println("Hello Private")
}

func SayHelloFromPrivate() {
	generateHello()
}
