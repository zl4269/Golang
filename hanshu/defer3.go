//探究defer函数的作用3：修改函数的具名返回值（返回值一定是要具名）
package main

import "fmt"

func foo(a, b int) (x, y int) { //返回值具名，不能写成下int int
	defer func() {
		x *= 5
		y *= 10
	}()

	x = a + 5
	y = b + 6
	return x, y //因为是具名，return 后面可跟可不跟返回值变量x y
}

/* //不是具名返回值变量就回报错
func foo(a, b int) (int, int) {
	defer func() {
		x *= 5
		y *= 10
	}()

	x := a + 5
	y := b + 6
	return x, y
} */

func main() {
	x, y := foo(1, 2)
	fmt.Printf("x= %d , y= %d \n", x, y)
}
