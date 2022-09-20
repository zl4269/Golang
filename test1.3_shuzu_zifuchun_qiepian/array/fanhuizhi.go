/* 知识点：返回值是函数类型的函数
*/
package main

import "fmt"

//函数的返回值是函数类型func()  //a函数的返回值类型是func()  //把函数作为返回值
func a() func() int {
	return func() int {
		a := 3
		fmt.Println("hello")
		return a
	}
}

func main() {
	ans := a() //ans是一个函数
	b := ans() //执行该函数
	fmt.Printf("Typeof b: %T\n", b)   //输出 int 
	fmt.Printf("Valueof b: %d\n", b)  //输出 3
}
