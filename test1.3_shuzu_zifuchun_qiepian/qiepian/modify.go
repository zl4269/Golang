package main

//测试切片的底层操作，底层是一个数组，切片就是打开的一个访问接口

import "fmt"

func main() {
	//切片cap的大小就是底层数组切片之后的元素
	var nums = [10]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	s1 := nums[1:5]
	s2 := nums[6:9]
	s3 := nums[3:7]
	fmt.Printf("lenofs1:%d   capofs1:%d    s1:%v\n", len(s1), cap(s1), s1)
	fmt.Printf("lenofs1:%d   capofs1:%d    s1:%v\n", len(s1), cap(s1), s1)
	fmt.Printf("lenofs2:%d   capofs2:%d    s2:%v\n", len(s2), cap(s2), s2)
	fmt.Printf("lenofs2:%d   capofs2:%d    s2:%v\n", len(s2), cap(s2), s2)
	fmt.Printf("lenofs3:%d   capofs3:%d    s3:%v\n", len(s3), cap(s3), s3)
	fmt.Printf("lenofs3:%d   capofs3:%d    s3:%v\n", len(s3), cap(s3), s3)

	//一个配片更改元素，输出其他切片的对应部分内容都会发生更改，因为更改了底层数组
	fmt.Println("############################################")
	s1[1] = 0
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)

}
