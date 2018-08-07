package main

import (
	"fmt"
	"reflect"
	"unsafe"
)


func main()  {
	//var buff bytes.Buffer
	a:= []byte{55, 101, 51, 53, 50, 57, 57, 52, 102, 49, 100 ,56 ,49, 51 ,99 ,57}
	fmt.Println("testfile.1",a)

	b := BytesToString(a)
	fmt.Println("bbbbb  ", b)
}


func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
