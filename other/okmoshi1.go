package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)
	c <- 1
	c <- 2
	c <- 3
	close(c) // 这里就关闭了channel,下面依然可以读取到channel里面残留的数据,直到数据处理完成后才读取到关闭消息
	// c <- 4	// 本行代码不可执行,向一个已关闭的通道压入数据时会造成panic

	// 所以证明了根据从channel读取到的第二个字段的值为true或者false时,并不能证明channel是否未关闭或已关闭
	// 总结
	// 1. 第二个字段的值为true时，channel可能没关闭，也可能已经关闭，不能证明什么
	// 2. 第二个字段为false时，可以证明channel中已没有残留数据且已关闭

	for {
		i, isOpen := <-c
		if !isOpen { // 若不加本判断,从一个已关闭的channel中读取数据会直接返回,且是默认值,造成死循环!!!!!!
			fmt.Println("channel 已关闭!")
			break
		}
		fmt.Printf("i=%v, isOpen=%v \n", i, isOpen)
	}
}
