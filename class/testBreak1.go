//设置break 和 continue跳出的位置 break[label]
//语法:
/*
outerLoop:
	for i:=0;i<n;i++{
		//....
		for j :=0;j<m;j++{
			//当不满足条件的时候，直接终止outerLoop代替的循环的执行（即最外层的for循环）
			break outerLoop

			//当条件满足某些条件的时候，直接跳出outerLoop所代替的循环，继续执行
			continue outerLoop
		}
	}
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan interface{})
	go func() {
	loop:
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("tick")
			case <-exit:
				fmt.Println("exiting....")
				break loop //loop就是设定标签的名字，可以任意取。 loop代替的就是for循环,break loop就是braek for循环
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
		exiting....
		exit!

	*/
}

//设置一个
