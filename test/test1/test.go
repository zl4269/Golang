//测试一下 <<  |= 操作符

package main

import "fmt"

var specialBttes [16]byte

func main() {
	for _, b := range []byte(`\.+*?()|[]{}^$`) {
		specialBttes[b%16] |= 1 << (b / 16) // |= 表示左右两边的数做与运算得到的结果赋值给左边的数  1<< n  表示乘2的n次方
		/* 		fmt.Printf("b: %d   ", b)
		   		fmt.Printf("b/6:%v   ", b/16)
		   		fmt.Println(b % 16) */
	}
	fmt.Printf("specialBttes: %v\n", specialBttes)
	//specialBttes: [0 0 0 0 4 0 0 0 4 4 4 164 160 160 36 8]  注意每个结果的由来，b%16可能是是数组的同一个元素

	fmt.Println("##################################")
	x := 64
	x >>= 4 //右移相当于除2的n次方，右移位数太多就可能输出0、
	fmt.Printf("x: %v\n", x)
}
