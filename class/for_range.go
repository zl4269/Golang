//掌握两点：
//1.for range中迭代变量i,v是重用的
//2.闭包函数：函数 + 引用环境
//3.for 循环的声明变量属于隐式代码块，作用域为整个循环

package main

import (
	"fmt"
	"time"
)

func demo1() {
	var m = [...]int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 10)
}

//输出结果也不一样，由goroutine的调用顺序有关
func demo2() {
	var m = [...]int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func(i, v int) { //将i v与goroutine的闭包函数绑定
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}(i, v) //传入参数
	}
	time.Sleep(time.Second * 10)
}

func main() {
	demo1()
	fmt.Println("#########################")
	demo2()

}

/*
输出结果:
4 5
4 5
4 5
4 5
4 5
#########################
4 5
3 4
2 3
1 2
0 1
*/
