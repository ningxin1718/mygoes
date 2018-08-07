package main

import "fmt"

func main() {
	asciiStr := "ABC"
	asciiBytes := []byte(asciiStr)
	fmt.Printf("OK: string=%v, bytes=%v\n", asciiStr, asciiBytes)
	fmt.Printf("OK: byte(A)=%v\n", asciiBytes[0])
}
// OK: string=ABC, bytes=[65 66 67]
// OK: byte(A)=65