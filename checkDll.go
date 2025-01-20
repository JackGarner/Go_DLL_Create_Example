package main

import "C"
import (
	"fmt"
	"syscall"
)

func main() {
	dll := syscall.MustLoadDLL("example_from_go.dll")
	proc := dll.MustFindProc("sayHello")
	proc.Call()
	fmt.Println("Function executed successfully.")
}
