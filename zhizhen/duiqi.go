package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	A byte
	B int16
	C int32
	D int64
}

func main() {
	t := T{}

	//输出每个成员变量所占空间
	fmt.Println(unsafe.Sizeof(t.A)) //1
	fmt.Println(unsafe.Sizeof(t.B)) //2
	fmt.Println(unsafe.Sizeof(t.C)) //4
	fmt.Println(unsafe.Sizeof(t.D)) //8

	//输出每个成员变量的内存偏移量(注意偏移量的计算，原则，充分利用内存空间，但是又得方便查找（一遍查找）)
	fmt.Println(unsafe.Offsetof(t.A)) //0
	fmt.Println(unsafe.Offsetof(t.B)) //2
	fmt.Println(unsafe.Offsetof(t.C)) //4
	fmt.Println(unsafe.Offsetof(t.D)) //8

	fmt.Println(unsafe.Sizeof(t)) //16
}
