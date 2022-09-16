package main

//switch语句的一些总结
/*
1.switch语句先是计算switch后面的表达式，然后与每个case做对比（从上到下从左到右进行对比）,判断执行哪一个case
2.case 后面有多个表达式，则依次计算每个表达式的值，分别与switch做对比，一旦匹配成功，该case后面的表达式就不会再计算
3.如果case执行体重存在fallthrough关键字，在执行完该case的时候就会将执行权直接交给下一个case，略过了下一个case表达式的判断。
*/

import "fmt"

func Expr(n int) int {
	fmt.Println(n)
	return n
}

func main() {
	switch Expr(2) {
	case Expr(1), Expr(2), Expr(3):
		fmt.Println("enter into case1")
		fallthrough
	case Expr(4):
		fmt.Println("enter into case2")
	}
	/*
	   输出：
	   2
	   1
	   2
	   enter into case1
	   enter into case2
	*/
}
