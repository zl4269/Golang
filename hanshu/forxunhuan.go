//探究闭包函数最容易在for循环中出现的问题

package main

import "fmt"

func main() {
	var funcSlice []func() //函数数组()
	for i := 0; i < 3; i++ {
		fmt.Println(&i) //同时对比一下地址
		funcSlice = append(funcSlice, func() {
			fmt.Println(i)
			fmt.Println(&i) //同时对比一下地址=====>地址都是一样的,
		})
	}
	for i := 0; i < 3; i++ { //循环调用函数数组
		funcSlice[i]()
	} // 3 3 3  调用的时候才会去使用外部变量的值

	//如何实现遍历=====》将for循环的迭代变量作为参数传入到闭包函数中去
	var funcSlice1 []func()
	for i := 0; i < 3; i++ {
		func(i int) { //再创建一个闭包函数，每次将i值传入到内部的闭包函数（i是更新的）
			funcSlice1 = append(funcSlice1, func() {
				fmt.Println(i)
			})
		}(i)
	}
	for j := 0; j < 3; j++ {
		funcSlice1[j]() // 0 1 2
	}
}
