//当结构体嵌入多个接口类型而且这些接口类型存在交集的时候，Go调用的顺序如下
//1)优先选择结构体自身实现的方法
//2)如果结构体自身没有实现，那么查找结构体中嵌入接口类型的方法集合中是否有该方法，如果有，则提升为该结构体方法
//3)如果结构体嵌入了多个类型且这些接口类型方法集合存在交集，那么Go编译器将会报错，除非结构体实现了交集中的所有方法

package main

import "fmt"

type Interface interface {
	M1()
	M2()
	M3()
}
type Interface1 interface {
	M1()
	M2()
	M4()
}

type T struct {
	Interface //T类型中包含了接口类型
	Interface1
}

// func main() {
// 	t := T{}
// 	t.M1()
// 	t.M2()
// 结果：
// ./MethodsInterface3_2.go:26:3: ambiguous selector t.M1  报错有歧义，模棱两可
// ./MethodsInterface3_2.go:27:3: ambiguous selector t.M2
// }

func (t T) M1() {
	fmt.Println("T`s M1")
}
func (T) M2() {
	fmt.Println("T`s M2")
}

func main() {
	t := T{} //注意结构体的初始化nil

	//调用的是自己方法
	t.M1()
	t.M2()
}

/*
T`s M1
T`s M2
*/
