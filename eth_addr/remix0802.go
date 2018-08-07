package main

import (
	"fmt"
	"strings"
	"encoding/hex"
	"reflect"
	"github.com/ethereum/go-ethereum/accounts/abi"

)

type unpackTest struct {
	def  string      // ABI definition JSON
	enc  string      // evm return data
	want interface{} // the expected output
	err  string      // empty or error if expected
}

func main() {
	var unpackTest = unpackTest{
		def:  `[{"type": "bytes"}]`,
		enc:  "000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000200100000000000000000000000000000000000000000000000000000000000000",
		want: [32]byte{},
		err:  "abi: cannot unmarshal []uint8 in to [32]uint8",
	}

	def := fmt.Sprintf(`[{ "name" : "method", "outputs": %s}]`, unpackTest.def)
	fmt.Println("def::",def)

	myabi, err := abi.JSON(strings.NewReader(def))
	if err != nil {
		fmt.Println("invalid ABI definition %s: %v", def, err)
	}

	fmt.Println("test.enc-->",unpackTest.enc)
	encb, err := hex.DecodeString(unpackTest.enc)
	if err != nil {
		fmt.Println("invalid hex: %s" + unpackTest.enc)
	}

	fmt.Println("encb: ",encb)

	outptr := reflect.New(reflect.TypeOf(unpackTest.want))
	err = myabi.Unpack(outptr.Interface(), "method", encb)

	//if err := unpackTest.checkError(err); err != nil {
	//	fmt.Println("test %d (%v) failed: %v", i, unpackTest.def, err)
	//	return
	//}
	fmt.Printf("unpackTest.want: %v\n",unpackTest.want)
	out := outptr.Elem().Interface()

	fmt.Printf("out: %v",out)

	if !reflect.DeepEqual(unpackTest.want, out) {
		fmt.Printf("failed: (%v)  expected %v, got %v", unpackTest.def, unpackTest.want, out)
	}
}