package main

import (
	"syscall/js"
)

func add(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(3)
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("addFun", js.FuncOf(add))
	<-done
}
