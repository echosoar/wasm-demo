package main

import (
	"syscall/js"
)

func changeObj(this js.Value, args []js.Value) interface{} {
	key := args[0].String();
	this.Set(key, args[1]);
	return nil;
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("changeObj", js.FuncOf(changeObj))
	<-done
}
