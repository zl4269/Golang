//从键盘和便准输入os.Stdin读取输入，最常用的办法就是使用fmt包提供的Scan和Sscan函数
/*
Scan，Scanf和Scanln读取os.Stdin;
Fscan, Fscanf和Fscanln从指定的io.Reader读取;
Sscan, Sscanf和Sscanln从参数字符串中读取。

Scan，Fscan，Sscan可以将输入中的换行符视为空格。
Scanln、Fscanln和Sscanln在换行符处停止扫描，并要求项目后面跟着换行符或EOF----------按行读取。
Scanf、Fscanf和Scanf根据格式字符串解析参数，类似于Printf。
*/

// package main

// import "fmt"

// var (
// 	firstName, lastName, s string
// 	i                      int
// 	f                      float32
// 	input                  = "56.12 / 5212 / Go"
// 	format                 = "%f / %d / %s"
// )

// func main() {
// 	fmt.Println("Please enter your full name: ")
// 	fmt.Scan(&firstName, &lastName, &i) //可以读取到下一行，将换行符视为空格
// 	fmt.Println(firstName, lastName, i)

// 	fmt.Scanf("%s %s %d %f", &firstName, &lastName, &i, &f) //设置了接收的格式,输入内容用空格隔开对应每一个变量,跟Scanln一样，不能读取到换行后的数据
// 	fmt.Println(firstName, lastName, i, f)

// 	fmt.Scanln(&firstName, &lastName, &i) //按行读取，每次读取一行,输入内容用空格隔开对应每一个变量
// 	fmt.Println(firstName, lastName, i)

// }

package main

import "fmt"

func main() {
	var n int
	var nums []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		nums = append(nums, tmp)
	}
	fmt.Printf("nums:%v\n", nums)
}
