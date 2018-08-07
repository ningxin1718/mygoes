package main

import (
	"fmt"
	"github.com/usechain/go-usechain/crypto"
	"github.com/usechain/go-usechain/common/hexutil"
	"github.com/usechain/go-usechain/common"
)

func main()  {



	address := "0x7b65b7C1be089252BC830C6D309e7b066d320e53"
	fmt.Println(common.FromHex(address))


	msg := crypto.Keccak256([]byte(address))
	msg2:=hexutil.Encode(msg)
	fmt.Println("msg--->",msg2)

	msg3 := crypto.Keccak256([]byte(address))
	fmt.Println("msg2------->",msg3)


}
