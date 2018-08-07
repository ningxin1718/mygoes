package main

import (
	"fmt"
	"strings"
)

func main()  {
	a:="asdf"
	fmt.Println(a[0:2])

	x:=[]string{a[0:3],a[0:2]}
	fmt.Println(x)

	b:="asdf&aaaa+dddd"
	pa := strings.Split(b, "&")
	fmt.Println(pa)
}
