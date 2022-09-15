//探究是底层是怎么样扩容的
//切片就是一个结构体：包含了切片起始元素的指针，切片的长度len和切片的最大容量cap
package main

import "fmt"

func main() {
	var s []int
	s = append(s, 11)
	fmt.Printf("len(s): %d cap(s):%d\n", len(s), cap(s))
	s = append(s, 12)
	fmt.Printf("len(s): %d cap(s):%d\n", len(s), cap(s))
	s = append(s, 13)
	fmt.Printf("len(s): %d cap(s):%d\n", len(s), cap(s))
	s = append(s, 14)
	fmt.Printf("len(s): %d cap(s):%d\n", len(s), cap(s))
	s = append(s, 15)
	fmt.Printf("len(s): %d cap(s):%d\n", len(s), cap(s))
	//len结果：1 2 3 4 5 
	//cap结果：1 2 4 4 8 按照一定的扩容算法进行扩容（扩容到4的时候，容量还没有使用完，append后的容量还是4）
	//切片append的原理：1.当前容量不够，按照扩容算法扩容，容量够，直接追加；2.扩容就是动态分配新的数组，然后将旧数组中的元素复制到里面；3.旧数组就会被垃圾回收掉
}
