//探究defer函数的第1中用法：拦截panic
//recover()函数是Go语言中的内建函数，可以让进入宕机(panic)流程中的goroutine恢复过来，
//recover仅在延迟函数defer中有效，在正常的执行过程中，调用recover会返回nil并且没有任何其他的效果
//如果当前的goroutine陷入panic,调用recover可以捕获到panic的输入值，并且恢复正常的执行。
package main

import "fmt"

func bar() {
	fmt.Println("rause a panic")
	panic(-1)
}

func foo() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recovered from a panic")
		}
	}()
	bar()
}

func main() {
	foo()
	fmt.Println("main exit normally")
}

/*
输出结果：
rause a panic
recovered from a panic
main exit normally
*/
