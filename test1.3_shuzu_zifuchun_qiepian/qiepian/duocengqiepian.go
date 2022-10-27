package main

import "fmt"

func main() {
	var sli = make([]int, 5)
	for k, _ := range sli {
		sli[k] = k + 1
	}
	fmt.Printf("sli: %v\n", sli)

	//多重切片
	s1 := sli[1:5]
	fmt.Printf("s1: %v\n", s1)
	s2 := s1[1:2]
	fmt.Printf("s2: %v\n", s2)

	fmt.Println("##################")
	s2[0] = 100
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("sli: %v\n", sli) //sli: [1 2 100 4 5]依然会改变底层的数据
}
