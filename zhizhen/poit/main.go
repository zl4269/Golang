//unsafe.pointer用于访问操作结构体的私有变量

package main

import (
	"fmt"
	"poit/p"
	"unsafe"
)

func main() {
	var v *p.V = new(p.V) //new是golang的内置方法，用来分配一段内存(会按类型的零值来清零)，并返回一个指针。
	var i *int32 = (*int32)(unsafe.Pointer(v))
	*i = int32(98)
	var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int32(0))))) //Sizeof计算数据类型所占的字节数，在C++中传递的是类型，在go中传递的是类型的任意值
	*j = int64(763)
	v.PutI()
	v.PutJ()

	var b []byte = []byte{'a', 'b', 'c'}
	var c *byte = &b[0]
	fmt.Println(*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1)))) //输出的是b的ASCII码，因为偏移了一个字节(+ unitptr(1))

	/*
		运行结果:
		size=16
		Alignof(w.b)=1
		Alignof(w.i)=4
		Alignof(w.j)=8
		Alignof(W{})=8
		i=98
		j=0
		98
		//分析：前四个结果在后两个结果之前是因为前四个结果是在初始化中的到的，Go语言在调用包的时候，会有优先执行其初始化函数main
		//size=16而不是等于13（1+4+8），里面涉及了Go语言中的内存对齐。
		为什么最大对齐是8,而其中的字节数只占了16(8+8)--->按照8对齐的,这个成员变量在结构体中的顺序有关。因为b i相邻，分配8个字节就能存储两个变量了。
		如果将j放在b i中间，那么size 就会等于24。结果如下:
		size=24
		Alignof(w.b)=1
		Alignof(w.i)=4
		Alignof(w.j)=8
		Alignof(W{})=8
		i=98
		j=0
		98
	*/

}
