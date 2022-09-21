//主要探究在for循环中闭包函数的工作原理

package main

import "fmt"

func main() {
	var dummy [3]int
	var f func() //函数类型
	for i := 0; i < len(dummy); i++ {
		f = func() { //每次循环都会更新f
			fmt.Println(i)
		}
	}
	f() //3

	//为什么是3不是2
	//上面函数可以写成一下函数
	var dummy1 [3]int
	var f1 func()
	for i := 0; i < len(dummy1); { //注意第二个条件后面的分号
		f1 = func() {
			fmt.Println(i)
		}
		i++
	}
	f1() //3  因为i=3时条件不成立退出循环，但此时i=3

	//使用for range循环代替for循环
	var dummy2 [3]int
	var f2 func()
	for i, _ := range dummy2 {
		f2 = func() {
			fmt.Println(i)
		}
	}
	f2() //2

}
