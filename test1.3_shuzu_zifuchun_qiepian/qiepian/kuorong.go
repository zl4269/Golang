//探究是底层是怎么样扩容的
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

}
