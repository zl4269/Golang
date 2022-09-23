//当结构体嵌入多个接口类型而且这些接口类型存在交集的时候，Go调用的顺序如下
//1)优先选择结构体自身实现的方法
//2)如果结构体自身没有实现，那么查找结构体中嵌入接口类型的方法集合中是否有该方法，如果有，则提升为该结构体方法
//3)如果结构体嵌入了多个类型且这些接口类型方法集合存在交集，那么Go编译器将会报错，除非结构体实现了交集中的所有方法

package main

import (
	"fmt"
)

type Interface interface {
	M1()
	M2()
}

type T struct {
	Interface //T类型中包含了接口类型
}

func (T) M1() { // 没有用到 t T中的t，可以省略。
	fmt.Println("T`s M1")
}

type S struct{}

func (S) M1() {
	fmt.Println("S`s M1")
}

func (S) M2() {
	fmt.Println("S`s M2")
}

func main() {
	var t = T{
		Interface: S{}, //注意：S结构体实现了Interface中的所有方法，解实现了Interface接口，所以S可以赋值给Interface结构的变量。
	}
	t.M1() //先调用自己的方法，有就直接调用，没有的话查找嵌入结构体的方法，然后将其提升为自己的方法。
	t.M2()
}

/*
T`s M1
S`s M2
*/
