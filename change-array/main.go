package main

import (
	"syscall/js"
)

func changeArray(this js.Value, args []js.Value) interface{} {
	arrayValue := js.Global().Get("arrayData");
	currentLength := arrayValue.Length();
	doType := args[0].String();
	switch doType {
	case "push":
		arrayValue.SetIndex(currentLength, args[1]);
	case "pop":
		newLength := currentLength - 1;
		popValue := arrayValue.Index(newLength);
		arrayValue.Set("length", newLength);
		return popValue;
	}
	return nil;
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("changeArray", js.FuncOf(changeArray))
	<-done
}
