//主要验证当切片触碰到数组上届的时候，切片就会跟原数组进行分离
//当切片append之后，其底层数组还有空间，就改变底层数组空间的值
//当切片append之后，底层数组没有空间的话，就会与原数组进行分离，创建自己的底层数组

package main

import "fmt"

func main() {
	u := []int{11, 12, 13, 14, 15}
	fmt.Printf("u: %v\n", u) //u: [11 12 13 14 15]

	s := u[1:3]
	fmt.Printf("slice(len=%d,cap=%d):%v\n", len(s), cap(s), s) //slice(len=2,cap=4):[12 13]   切片

	s = append(s, 24)
	fmt.Printf("after append 24, u: %v\n", u)                                   //after append 24, u: [11 12 13 24 15]  切片会改变底层数组
	fmt.Printf("after append 24, slice(len=%d,cap=%d):%v\n", len(s), cap(s), s) //after append 24, slice(len=3,cap=4):[12 13 24]

	s = append(s, 25)
	fmt.Printf("after append 25, u: %v\n", u)                                   //after append 25, u: [11 12 13 24 25]
	fmt.Printf("after append 25, slice(len=%d,cap=%d):%v\n", len(s), cap(s), s) //after append 25, slice(len=4,cap=4):[12 13 24 25]

	s = append(s, 26)
	fmt.Printf("after append 26, u: %v\n", u)                                   //after append 26, u: [11 12 13 24 25]   //底层数组u没有发生改变，说明分离了
	fmt.Printf("after append 26, slice(len=%d,cap=%d):%v\n", len(s), cap(s), s) //after append 26, slice(len=5,cap=8):[12 13 24 25 26]

	//验证是否真的分离了
	s[0] = 22
	fmt.Printf("after reassign 1st elem of slice, u: %v\n", u)                                   //after reassign 1st elem of slice, u: [11 12 13 24 25]  //u数组没有发生改变说明确实分离了
	fmt.Printf("after reassign 1st elem of slice, slice(len=%d,cap=%d):%v\n", len(s), cap(s), s) //after reassign 1st elem of slice, slice(len=5,cap=8):[22 13 24 25 26]

}
