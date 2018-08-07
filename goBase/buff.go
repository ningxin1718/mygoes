package main

import (
	"bytes"
	"fmt"
)

func main()  {
	var buff bytes.Buffer
	a := []byte{1,2,3}
	buff.Write(a)
	fmt.Println(a)

	b := buff.Bytes()

	fmt.Println(string(b))

	rest:= string(buff.Bytes())
	fmt.Println(rest)

}
