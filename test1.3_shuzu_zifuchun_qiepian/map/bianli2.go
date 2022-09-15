//利用切片记录map的下标，从而做到有序遍历
//下面的结果是一样的，但是所有结果的顺序是随机的，稍微改动也能实现遍历的顺序就是插入的顺序
package main

import "fmt"

func doIteration(s []int, m map[int]int) {
	fmt.Printf("{ ")
	for _, k := range s {
		v, ok := m[k]
		if !ok {
			continue
		}
		fmt.Printf("[%d %d]", k, v)
	}
	fmt.Printf("} \n")
}

func main() {
	var s []int
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}
	for k, _ := range m {
		s = append(s, k)
	}

	for i := 0; i < 3; i++ {
		doIteration(s, m) //结果是一样的
	}

}
