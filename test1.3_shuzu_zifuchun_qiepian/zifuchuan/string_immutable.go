package main

//尝试强制改变字符串的底层结构来挑战字符串的只读性（会报错）
import (
	"fmt"
	"unsafe"
)

func modifyString(s *string) {
	//取出第一个8字节的值
	p := (*uintptr)(unsafe.Pointer(s))

	//获取底层数组的地址
	var array *[5]byte = (*[5]byte)(unsafe.Pointer(*p))

	var len *int = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Sizeof((*uintptr)(nil))))

	for i := 0; i < (*len); i++ {
		fmt.Printf("%p => %c\n", &((*array)[i]), (*array)[i])
		p1 := &((*array)[i])
		v := (*p1)
		(*p1) = v + 1 //尝试改变底层元素的值
	}
} 

func main() {
	//原始字符串
	var s string = "hello"
	fmt.Printf("Befer change s: %v\n", s)

	//通过unsafe指针改变原始string
	modifyString(&s)
	fmt.Printf("After change s: %v\n", s)

}
