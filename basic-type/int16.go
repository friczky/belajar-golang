package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a int16 = 1111

	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
