package main

import (
	"go-wanchain-feature-scan/common/hexutil"
	"fmt"
	"encoding/hex"
	"bytes"
)

func main()  {
	x:="0x3565333532393934663164383133633939653262663431343337366333636464623865343030633564663536306639303366666262356338636135393135643732323766393662356263303464653137323133323733393631336439653134333363643264666564643062666234393031316266643231333239376231343861"

	x1,_:=hexutil.Decode(x)

	fmt.Println("x1:::",x1)
	fmt.Println("string(x1): ",string(x1))

	a:= []byte{53, 101, 51, 53, 50, 57, 57, 52, 102, 49, 100 ,56 ,49, 51 ,99 ,57 ,57 ,101 ,50 ,98, 102, 52 ,49 ,52, 51, 55, 54, 99, 51 ,99, 100 ,100}
	fmt.Println("hex.EncodeToString(a): ",hex.EncodeToString(a))
	fmt.Println("string(a):",string(a))

	b:=[]byte{98,56,101,52,48, 48, 99}
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	fmt.Println("buffer: ", buffer)
	buffer.Write(a)
	buffer.Write(b)
	b3 :=buffer.Bytes()
	fmt.Println("b3: ",b3)
	fmt.Println("string(b3): ",buffer.String())
}
