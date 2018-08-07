package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"encoding/hex"
)

func main() {
	buf := new(bytes.Buffer)

	var source uint32 = 15
	//source := []uint32{15}
	err := binary.Write(buf, binary.LittleEndian, source)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	fmt.Println(buf.Bytes())
	fmt.Println("::::::",hex.EncodeToString(buf.Bytes()))


	check := make([]uint32, 3)
	rbuf := bytes.NewReader(buf.Bytes())
	err = binary.Read(rbuf, binary.LittleEndian, &check)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Printf("Decoded: %v\n", check)

}

