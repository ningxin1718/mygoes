package main

import (
	"fmt"

	"strconv"
	"encoding/hex"
)

func main()  {
	stringInt()
}

func stringInt()  {
	x:=strconv.FormatInt(1000,16)

	method:="c03c1796"
	a:="0000000000000000000000000000000000000000000000000000000000000001"
	b:="0000000000000000000000000000000000000000000000000000000000000001"
	fmt.Println(a[0:len(a)-len(x)])
	fmt.Println(b)
	dataString := method +a[0:len(a)-len(x)]+x+ b
	fmt.Println(dataString)

	dst,err:=hex.DecodeString(dataString)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(dst)
}
