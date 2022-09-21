//理解函数变量
/*在 Go 语言中，函数被看作是第一类值，这意味着函数像变量一样，有类型、有值，
其他普通变量能做的事它也可以。*/

package main

import "fmt"

func mysquare(x int) {
	fmt.Println(x * x)
}

func main() {
	//直接调用
	num := 2
	mysquare(num)

	//函数变量：将函数当作一种数据类型去使用
	//函数类型定义的变量的使用方法跟其他变量的使用是一样的，唯一区别就是还可以当作函数去调用

	f := mysquare //注意不能使用括号
	if f != nil {
		fmt.Println("函数变量的零值是nil，并且函数变量可以跟nil进行比较") //注意函数变量不能跟函数变量进行比较
	}
	f(3) //使用该函数变量

}
