package main

import (
	"syscall/js"
)

func invokeClassRun(this js.Value, args []js.Value) interface{} {
	jsClass := js.Global().Get("classA");
	instance := jsClass.New(123);
	result := instance.Call("run");
	return js.ValueOf(result);
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("invokeClassRun", js.FuncOf(invokeClassRun))
	<-done
}
