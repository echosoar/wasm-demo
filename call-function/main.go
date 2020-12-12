package main

import (
	"syscall/js"
)

func callFun(this js.Value, args []js.Value) interface{} {
	funcName := args[0].String();
	fun := js.Global().Get(funcName);
	result := fun.Invoke(args[1]);
	return js.ValueOf(result)
}

func callFun2(this js.Value, args []js.Value) interface{} {
	funcName := args[0].String();
	result := js.Global().Call(funcName, args[1]);
	return js.ValueOf(result)
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("callFun", js.FuncOf(callFun))
	js.Global().Set("callFun2", js.FuncOf(callFun2))
	<-done
}
