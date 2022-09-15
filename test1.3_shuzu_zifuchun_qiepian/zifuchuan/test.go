package main

import "fmt"

type buf [32]byte

func main() {
	var b *buf
	*b = buf{}
	fmt.Println(b[0])
}
