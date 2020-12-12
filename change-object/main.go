package main

import (
	"syscall/js"
)

func changeObj(this js.Value, args []js.Value) interface{} {
	key := args[0].String();
	doType := args[1].String();
	switch doType {
	case "set":
		this.Set(key, args[2]);
	case "delete":
		this.Delete(key);
	}
	return nil;
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("changeObj", js.FuncOf(changeObj))
	<-done
}
