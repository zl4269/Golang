//介绍遍历map---》无序遍历的，并且介绍怎么有序遍历
package main

import "fmt"

func doIteration(m map[int]int) {
	fmt.Printf("{ ")
	for k, v := range m {
		fmt.Printf("[%d %d]", k, v)
	}
	fmt.Printf("} \n")
}

func main() {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}
	for i := 0; i < 3; i++ {
		doIteration(m) //每次输出的结果都不一样，因为迭代器对初始位置做了随机处理
	}
}
