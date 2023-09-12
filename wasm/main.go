// Note for Golang users: edit settings -> change build and arch to the following for type safety
//go:build js && wasm

package main

import (
	"syscall/js"
)

func myGolangFunction(this js.Value, args []js.Value) any {
	return args[0]

}

func main() {
	println("Hello, WebAssembly!")
	js.Global().Set("myGolangFunction", js.FuncOf(myGolangFunction))
	<-make(chan bool)
}
