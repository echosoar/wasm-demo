package main

import (
	"syscall/js"
)

func typeTransform(this js.Value, args []js.Value) interface{} {
	newArgs := make([]interface{}, len(args))
	for index, arg := range args {
		newArgs[index] = transform(arg);
	}
	return js.ValueOf(newArgs);
}

func transform(v js.Value) interface{} {
	typeStr := v.Type().String();
		switch typeStr {
		case "number":
			return v.Float();
		case "string":
			return v.String();
		case "boolean":
			if v.Truthy() {
				return true;
			} else {
				return false;
			}
		case "null":
			return js.Null();
		case "undefined":
			return js.Undefined();
		case "function":
			return v.Invoke();
		case "object":
			arrayConstructor := js.Global().Get("Array");
			if v.InstanceOf(arrayConstructor) {
				arrayLen := v.Length();
				newArray := make([]interface{}, arrayLen);
				for i:=0;i<arrayLen;i++ {
					newArray[i] = transform(v.Index(i));
				}
				return newArray;
			}

			objectConstructor := js.Global().Get("Object");

			if v.InstanceOf(objectConstructor) {
				allKeys := js.Global().Get("Object").Call("keys", v);
				allKeysLen := allKeys.Length();
				newMap := make(map[string]interface{});
				for i:=0; i< allKeysLen;i++ {
					key := allKeys.Index(i).String();
					newMap[key] = transform(v.Get(key));
				}
				return newMap;
			}
			return typeStr;
		default:
			return typeStr;
		}
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("typeTransform", js.FuncOf(typeTransform))
	<-done
}
