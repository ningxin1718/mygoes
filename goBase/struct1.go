package main

import "fmt"

type storj struct {
	index 		int
	parameter 	int
	methodType	string
}

var (
	OneTimeAddrConfirmedLenIndex = storj{
		index:3,
		parameter:0,
		methodType:"uintValue",
	}
)


func main()  {

	var a string


	fmt.Println(OneTimeAddrConfirmedLenIndex)
	fmt.Println(a)


}
