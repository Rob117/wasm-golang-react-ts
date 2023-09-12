// Note for Golang users: edit settings -> change build and arch to the following for type safety
// https://stackoverflow.com/a/67020638/4982995
//go:build js && wasm

package main

import (
	"syscall/js"
)

func myGolangFunction(_ js.Value, args []js.Value) any {
	return args[0].Int() + args[1].Int()
}

func main() {
	println("Hello, WebAssembly!")
	js.Global().Set("myGolangFunction", js.FuncOf(myGolangFunction))
	<-make(chan bool)
}
