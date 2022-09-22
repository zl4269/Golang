//使用defer关键字的关键在于把握好后面表达式的求值时期
//牢记：defer关键字后面的表达式是在将deferred函数注册的时候进行求值的

package main

import "fmt"

func foo1() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i) //注册的时候，依次注册fmt.Println(0)  ;mt.Println(1)  ;mt.Println(2)  ;mt.Println(3)  ;  最后调用的时候先进后出，结果为 3 2 1 0
	}

}

func foo2() {
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i) //注册的时候传入了参数，所以也是fmt.Println(0)  ;mt.Println(1)  ;mt.Println(2)  ;mt.Println(3)  ;  最后调用的时候先进后出，结果为 3 2 1 0
	}
}

func foo3() {
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}() //注册的时候没有传入了参数，即fmt.Println()  ;mt.Println()  ;mt.Println()  ;mt.Println()  ;  最后调用的时候以闭包的形式调用i先进后出，结果为 4 4 4 4
	}
}

func main() {
	fmt.Println("The result foo1 :")
	foo1()

	fmt.Println("The result foo2 :")
	foo2()

	fmt.Println("The result foo3 :")
	foo3()

}

/*
结果：
The result foo1 :
3
2
1
0
The result foo2 :
3
2
1
0
The result foo3 :
4
4
4
4
*/
