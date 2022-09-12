package main

import (
	"fmt"
	"unicode/utf8"
)

func RunCount(b []byte) int { //统计[]byte中所占的字符数量（[]byte中所占的字节数已知，但是其具体有多少个字符就是这样统计的）
	count := 0
	for i := 0; i < len(b); {
		if b[i] < utf8.RuneSelf { //RuneSelf  = 0x80（0x80=128,<0x80表示的是占单个字节的字符）
			i++
		} else {
			v, n := utf8.DecodeRune(b[i:]) //n最大为4，汉字占3个字节
			fmt.Printf("v: %v\n", v)       //DecodeRune返回值：字符对应的Unicode编码值和其占内存的字节数
			i += n
		}
		count++
	}
	return count
}

func main() {
	str := "hello, 世界"
	ans := RunCount([]byte(str))
	fmt.Println(ans) //输出9，空格也是一个字符
}
