package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan interface{})

	go func() {
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("tick")
			case <-exit:
				fmt.Println("exiting....")
				break //只会跳出select循环
			}
		}
		fmt.Println("exit!")
	}()
	time.Sleep(3 * time.Second)
	exit <- struct{}{} //定义一个结构之后使用他的零值

	//等待协程退出
	time.Sleep(3 * time.Second)
	/*
	   输出结果：
	   tick
	   tick
	   exiting....  //直接break出select循环，但是没有跳出for循环
	   tick
	   tick   //终止是因为主函数结束了
	*/
}

//设置一个
