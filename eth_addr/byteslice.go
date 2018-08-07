//@Time  : 2018/3/29 14:27
//@Author: Greg Li
package main

import (
	"fmt"


)

type CommitteePubs []byte

func main() {

	var tmp CommitteePubs
	fmt.Println(tmp)
	var committee  = make([]string, 3)
	fmt.Println(committee)
	fmt.Println(len(committee))

}

