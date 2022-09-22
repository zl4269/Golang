//介绍延续传递CPS
//CPS风格:函数不允许返回值。
//一个函数A将其想返回的值显示的传递给一个函数continuation,而这个continuation函数自身是函数A的一个参数。
//举例：求阶乘-----这个例子虽然是反例，但是比较经典
package main

import "fmt"

//用递归实现
func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

//现在采用CPS风格。
//1.去掉函数中的返回值，为其添加一个函数类型的参数f
//2.将返回结果传递给函数f。即把return语句换成对f函数的调用（continuation）,f的函数体就是输出
func factorialcps(n int, f func(int)) {
	if n == 1 {
		f(n)
	} else {
		factorialcps(n-1, func(y int) { f(n * y) }) //把未完成的计算过程封装成一个新函数符f`,作为递归调用的第二个参数。
		//f`的参数y就是原factorialcps(n-1)的计算结果，而n*y就是要传递给函数f的（相当于返回值）
		//所以f`就定义为func(y int){f(n*y)}
	}
}

func main() {
	res := factorial(6)
	fmt.Printf("res: %v\n", res)
	factorialcps(6, func(y int) { fmt.Printf("%d\n", y) })

}
