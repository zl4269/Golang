/* 知识点:闭包
闭包 = 函数 + 引用环境
出现在函数内部有定义函数的情况,内部函数跟内部函数所处的环境构成闭包
注意内部函数不是把环境中的变量当成参数传递的，而是直接搜索
*/

package main

import "fmt"

func a() func() int {
	num := 3
	return func() int {
		return num
	}

}

func main() {
	ans := a()
	b := ans()
	fmt.Printf("b:%v\n", b)
}
