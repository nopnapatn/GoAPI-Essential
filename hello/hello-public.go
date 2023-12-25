package hello

import (
	"fmt"

	"github.com/nopnapatn/go-module/hello/internal/he"
)

func SayHelloFromPublic() {
	fmt.Println("Hello Public")

	// can call function public/private form another file with the same package
	generateHello()

	// can call function public/private from internal folder
	he.SayHe()
}
