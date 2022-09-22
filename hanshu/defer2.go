//探究defer的作用2---输出调试信息（在资源申请之后被注册，在资源释放时被调用）
//deferred函数被注册及调度执行的使时间点使得它十分适合用来输出一些调试信息
package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a")) //注册的时候已经把trace()函数执行了
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()
}

/*
输出结果：
entering: b
in b
entering: a
in a
leaving: a
leaving: b
*/
