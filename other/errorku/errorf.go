//介绍fmt.Errorf()函数的使用
//函数的声明：func Errorf(format string, a ...interface{}) error

package main

import (
	"fmt"
	"time"
)

func main() {
	//将fmt.Printf()中的内容保存在err中
	err := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("An error happened:", err)

	const name, dept = "GeeksforGeeks", "CS"
	err1 := fmt.Errorf("%s is a %s Portal.", name, dept)
	//fmt.Println(err1.Error())
	fmt.Println(err1)
}
