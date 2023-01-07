package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{})
	registerCallbacks()
	<-c
}

func HalloWasm(this js.Value, args []js.Value) interface{} {
	println("Hello  Wasm!!")
	return nil
}

func registerCallbacks() {
	js.Global().Set("HelloWasm", js.FuncOf(HalloWasm))
}