package main

import (
	"fmt"
	"encoding/hex"
)

func main()  {
	b:="0000000000000000000000000000000000000000000000000000000000000001"
	fmt.Println(b)


	a:="abc"
	k:=hex.EncodeToString([]byte(a))
	fmt.Println(k)

	fmt.Println(string([]byte(a)))

}

