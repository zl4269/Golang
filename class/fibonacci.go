package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case <-quit: //注意：是冒号//由于容量为0，没有输入就会一直阻塞
			fmt.Println("quit")
			return //返回值为空  //输出qiut之后就返回

		case c <- x:
			x, y = y, x+y
		}
	}

}

func main() {
	c := make(chan int)
	quit := make(chan int)

	//这个协程用于输出计算的结果(输出10个结果)
	//一定要在调用函数之前开启此协程，否者在执行输出c<-x的时候会发生阻塞从而死锁
	go func() {
		for i := 0; i < 10; i++ { //输出斐波那契数列的前十个
			fmt.Println(<-c)
		}
		//程序终止条件
		quit <- 0
	}()
	fibonacci(c, quit)
}
