package main

import "fmt"

func main() {
	var times [5][0]int
	fmt.Printf("times:%v\n", times)

	//忽略了迭代时的下标，相当于计算行数
	for range times {
		fmt.Println("hello")
	}

	//不忽略下标
	for i := range times {
		fmt.Printf("第%d个Hello\n", i)
	}

}
