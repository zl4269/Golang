package main

import "fmt"

func main() {
	//长度为0的数组(用于通道的同步操作-----》只是用于消息的同步，但是有没有额外分配空间)
	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	b := <-c1
	fmt.Printf("Type of b:%T\n", b)  //输出[0]int
	fmt.Printf("Value of b:%v\n", b) //输出:[]

	//一般更倾向于使用无类型名的匿名结构体代替空数组main
	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{}
	}()
	d := <-c2
	fmt.Printf("Type of d:%T\n", d)  //输出struct{}
	fmt.Printf("Value of d:%v\n", d) //输出:{}   {}表示是一个结构体

}
