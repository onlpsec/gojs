package main

import (
	"fmt"

	"github.com/dop251/goja"
)

const SCRIPT = `
function sum(a, b) {
    return +a + b;
}
`

func main() {

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
}
