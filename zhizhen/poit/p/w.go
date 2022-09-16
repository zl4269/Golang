package p

import (
	"fmt"
	"unsafe"
)

type W struct {
	b byte
	i int32
	j int64
}

func init() {
	var w *W = new(W)
	fmt.Printf("size=%d\n", unsafe.Sizeof(*w))
	fmt.Printf("Alignof(w.b)=%d\n", unsafe.Alignof(w.b)) //输出每个成员变量在w中的对齐值
	fmt.Printf("Alignof(w.i)=%d\n", unsafe.Alignof(w.i))
	fmt.Printf("Alignof(w.j)=%d\n", unsafe.Alignof(w.j))
	fmt.Printf("Alignof(W{})=%d\n", unsafe.Alignof(W{}))
}

/* type W struct {
	b byte
	j int64
	i int32
}

func init() {
	var w *W = new(W)
	fmt.Printf("size=%d\n", unsafe.Sizeof(*w))
	fmt.Printf("Alignof(w.b)=%d\n", unsafe.Alignof(w.b)) //输出每个成员变量在w中的对齐值
	fmt.Printf("Alignof(w.i)=%d\n", unsafe.Alignof(w.i))
	fmt.Printf("Alignof(w.j)=%d\n", unsafe.Alignof(w.j))
	fmt.Printf("Alignof(W{})=%d\n", unsafe.Alignof(W{}))
} */
