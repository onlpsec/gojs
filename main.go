package main

import (
	"fmt"
	"syscall/js"

	"github.com/onlpsec/gojs/goja"
)

const SCRIPT = `
function sum(a, b) {
    return +a + b;
}
`

func test(this js.Value, inputs []js.Value) interface{} {
	vm := goja.New()
	_, err := vm.RunString(SCRIPT)
	if err != nil {
		panic(err)
	}

	var sum func(int, int) int
	err = vm.ExportTo(vm.Get("sum"), &sum)
	if err != nil {
		panic(err)
	}

	fmt.Println(sum(40, 2)) // note, _this_ value in the function will be undefined.
	// Output: 42

	return nil
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("test", js.FuncOf(test))
	<-c
}
