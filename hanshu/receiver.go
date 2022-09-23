//receiver接受的类型是T,或者是 *T
/*
对于T来说，采用的是复制传递，因此不会影响到原型，只会影响到副本
对于*T 来说，传递的是指向原本的地址,因此会改变原型。
但是不管是T类型还是*T类型，其定义的实例的使用方法都是一样的（即*T定义的实例不需要解引用）
*/
package main

import "fmt"

type T struct {
	a int
}

func (t T) M1() {
	t.a = 10
}

func (t *T) M2() {
	t.a = 11 //无论是T类型实例还是*T类型实例，都既可以调用receiver为T类型的方法，也可以调用receiver为*T类型的方法。
}

func main() {
	var t T
	fmt.Printf("t.a: %v\n", t.a) // 0  结构体零值可用

	t.M1()
	fmt.Printf("M1t.a: %v\n", t.a) //0  不会改变原型，只会改变副本

	t.M2()
	fmt.Printf("M2t.a: %v\n", t.a) //11  *T改变的是原型

}
