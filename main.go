package main

import "C"
import "fmt"

//export sayHello
func sayHello() {
	fmt.Println("Hello world!")
}

func main() {}
