//验证闭包函数是以引用的形式去调用外层变量的
package main

import "fmt"

func main() {
	//证明外部变量的生命周期是在闭包函数调用完之后
	x := 1
	f := func() { //后面没有接括号说明不是函数调用而是闭包函数
		fmt.Println(x)
	}
	x = 2
	x = 3
	f() // 3  调用闭包函数

	//区别
	y := 1
	func() { //直接调用闭包函数
		fmt.Println(y) // 1 调用完之后，闭包函数中变量的生命周期已经结束
	}()
	y = 2
	y = 3

	//上述案列可以写成这样
	z := 1
	fz := func() {
		fmt.Println(z) //1
	}
	fz()
	z = 2
	z = 3

	//证明闭包函数对外部参数的调用是引用形式main
	num := 1
	func() {
		fmt.Println(&num)
	}()
	fmt.Println(&num)
}
