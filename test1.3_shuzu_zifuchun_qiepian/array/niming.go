/* 知识点:
//匿名函数的定义--->区别普通函数就是没有函数名字
  func (参数)(返回值){
	函数体
  }
//匿名函数在}后面加上(参数)就能立刻执行  举例如下： */

package main

import "fmt"

func main() {

	//匿名函数(直接执行)
	func(x, y int) {
		fmt.Printf("%d\n", x+y)
	}(10, 20)

	//匿名函数(调用执行)

	result := func(x, y int) int {
		return x * y
	}(10, 20)
	fmt.Printf("%d\n", result)

}
