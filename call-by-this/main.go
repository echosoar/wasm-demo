package main

import (
	"syscall/js"
)

func callByThis(this js.Value, args []js.Value) interface{} {
	key := args[0].String();
	return js.ValueOf(this.Get(key));
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("callByThis", js.FuncOf(callByThis))
	<-done
}
