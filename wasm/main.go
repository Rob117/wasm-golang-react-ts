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
