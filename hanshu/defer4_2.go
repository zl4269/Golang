//deferred在传入指针或者地址的时候不会立刻解引用，注册的时候依赖是按指针或者地址去注册的
package main

import "fmt"

func foo1() {
	s1 := []int{1, 2, 3}
	defer func(a []int) {
		fmt.Println(a)
	}(s1)
	s1 = []int{3, 2, 1}
	_ = s1 //初始化了不想使用就直接匿名  结果为：[]int{1,2,3}
}

func foo2() {
	s2 := []int{1, 2, 3}
	defer func(p *[]int) { //注册的时候是按照地址注册的，即fmt.Println(&s2)  不会解引用求出切片的值
		fmt.Println(*p) //结果为： []int{3,2,1}
	}(&s2)
	s2 = []int{3, 2, 1}
	_ = s2
}

func main() {
	fmt.Println("The result of foo1: ")
	foo1()

	fmt.Println("The result of foo2")
	foo2()
}
