//区分字符串底层的两个数组：[]rune和[]byte
//字符串最底层是[]byte进行存储的，然后[]rune相当于为了更好存储其他符号而封装的0-4字节的unicode编码的数组
package main

import "fmt"

func main() {
	fmt.Printf("原始符号（汉字）====>Uincode码点值=====>UTF8编码\n")
	s := "中国欢迎您"

	rs := []rune(s) //记录的是码点值
	sl := []byte(s) //底层string存储

	for i, v := range rs {
		var utf8Bytes []byte
		for j := i * 3; j < 3*(i+1); j++ {
			utf8Bytes = append(utf8Bytes, sl[j])
		}
		fmt.Printf("%s =================>%X==============>%X\n", string(v), v, utf8Bytes) //%X表示大写字母表示
	}

}
