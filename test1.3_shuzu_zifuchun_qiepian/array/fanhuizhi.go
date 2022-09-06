/* 知识点：
//闭包函数：闭包 = 函数 + 引用的环境
*/
package main

import "fmt"

//函数的返回值是函数类型func()  //a函数的返回值类型是func()
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
	fmt.Printf("Typeof b: %T\n", b)
	fmt.Printf("Valueof b: %d\n", b)

}
