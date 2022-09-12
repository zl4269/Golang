/*
知识点:
1.Golang中的字符串有两种，uint8 (byte)代表的ASCII的一个字符，rune 代表一个 UTF-8 字符
当需要处理中文、日文或者其他复合字符时，则需要用到rune类型，rune实际是一个int32.rune是Go语言中一种特殊的数据类型，他是int32的别名
汉字采用utf-8编码，占用三个字节，编码后的值是int类型。字母采用ASCII编码。

2.重点：
for采用byte类型循环，for range采用rune类型循环

3.一般有特殊字符（汉字和其他字符）最好是转换成rune，只有数字和字母也可转换成byte。

4.无论是那种转换都会重新分配内存，并赋值字节数组------》这使得字符串在转变的过程中能够改变其值（解除了其只读的语义）

*/

package main

import "fmt"

func main() {
	//一.测试for循环----》采用的是byte类型
	var str = "hello 你好"
	fmt.Println("len(str):", len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()
	/*
		结果:
		len(str): 12  ————————》由于汉字是占3个字节，所以长度是12.len()是底层统计数组元素个数的，因为是byte，所以统计的是字节个数
		hello ä½ å¥½
	*/

	//二、测试for range 循环 ----->采用的是rune类型
	var str1 = "hello 你好"
	fmt.Println("len(str1):", len(str1))
	for _, v := range str1 {
		fmt.Printf("%c", v)
	}
	fmt.Println()
	/*
		结果:
		len(str1): 12
		hello 你好
	*/

	//三、修改字符串的内容------》首先需要将字符串转换成rune或者byte类型，修改完成之后在转成成string类型

	var s1 = "hello"
	byteStr := []byte(s1)
	byteStr[0] = 'H'
	fmt.Printf("byteStr[0]:%v\n", byteStr[0]) //输出：byteStr[0]:72 即输出的是ASCII码
	fmt.Println(string(byteStr))              //输出结果：Hello

	var s2 = "你好吗"
	runeS := []rune(s2)
	runeS[2] = '啊'
	fmt.Printf("runeS[2]:%v\n", runeS[2]) //输出：runeS[2]:21834  即输出的是Unicode编码值
	fmt.Println(string(runeS))            //输出结果：你好啊

}
