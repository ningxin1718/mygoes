package main

import (
	"fmt"
	"go-wanchain-feature-scan/common/hexutil"
)

func main() {
	msg1 := "0x0E806648E4a4355737857819D10Af858bf5a4aE9"
	msg2 := "0x0e806648e4a4355737857819d10af858bf5a4ae9"
	fmt.Println(hexutil.Decode(msg1))
	fmt.Println(hexutil.Decode(msg2))
	fmt.Println([]byte(msg1))
	fmt.Println([]byte(msg2))
}
