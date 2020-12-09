package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("thisIsOneTwoThree", 123)
}
