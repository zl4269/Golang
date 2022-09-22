//函数类型显示转换的底层要求：函数的原型是一致的
/*
即  函数类型1 = 函数类型2（待转换的函数类型3）
将函数类型3装换成函数类型2，赋值给函数类型一，
如果不能将函数类型3直接转换成函数类型1，那就借助函数类型2--->前提是函数类型2必需实现了函数类型1的接口
（在Go语言中只要任何类型实现了接口的方法就相当于实现了该接口）
*/

package main

import "fmt"

//Go语言提供了接口，接口就是把所有具有共性的方法定义在一起
//接口相当于一种数据类型，任何其他类型只要实现了这些方法（全部实现）就是实现了这个接口。
type BinaryAdder interface { //定义接口，只有一个方法
	Add(int, int) int
}

type MyAdderFunc func(int, int) int //定义类型

func (f MyAdderFunc) Add(x, y int) int { //MyAdderFunc类型去实现Add方法就是实现了BinaryAdder接口
	return f(x, y) //函数类型的调用
}

func Myadd(x, y int) int {
	return x + y
}

func main() {
	var i BinaryAdder = MyAdderFunc(Myadd) //不能直接将Myadd直接赋值给i，因为他们的原型不是一样的
	fmt.Println(i.Add(5, 6))

}
