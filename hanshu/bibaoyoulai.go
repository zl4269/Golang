//函数作为函数的返回值说明闭包的由来
//定义：闭包函数是在函数内部定义的匿名函数，并且允许匿名函数访问定义他的外部函数的作用域。
package main

import "fmt"

func incr() func() int { //函数的返回值类型也是一个函数
	var x int
	return func() int {
		x++
		return x
	}

}

func main() {
	i := incr() //通过把这个函数返回值赋值给 i，i 就成为了一个闭包（incr()就是一个闭包），i 保存着对 x 的引用，可以想象 i 中有着一个指针指向 x 或 i 中有 x 的地址。

	// i 有着指向 x 的指针，所以可以修改 x，且保持着状态：也就是说，x 逃逸了，它的生命周期没有随着它的作用域结束而结束。
	//这里调用一次闭包三次
	fmt.Println(i()) //1
	fmt.Println(i()) //2
	fmt.Println(i()) //3

	//如果是直接调用该函数，相当于调用该闭包一次
	//这里调用了三次 incr()，返回了三个闭包，这三个闭包引用着三个不同的 x，它们的状态是各自独立的。
	fmt.Println(incr()()) //1  //incr()是调用incr函数，再加一个括号是调用其返回值函数
	fmt.Println(incr()()) //1
	fmt.Println(incr()()) //1

}
