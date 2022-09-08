package main

import "fmt"

func main() {
	//字符串的切片
	s := "hello, world"
	hello := s[:5]
	world := s[7:]

	fmt.Printf("############################1\n")
	fmt.Printf("hello:%v\n", hello)
	fmt.Printf("world:%v\n", world)

	fmt.Printf("############################2\n")
	//上述中的s可以直接使用字符串表示
	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Printf("s1:%v\n", s1)
	fmt.Printf("s2:%v\n", s2)

	//不能用for range遍历字符串，使用for range只能遍历字符串UTF8解码之后的Uni code码点值
	//输出的是Unicode字符，0是空字符，\uFFFD表示错误字符
	for _, v := range s {
		fmt.Printf("v:%#v\n", v)
	}
	fmt.Printf("############################3\n")
	fmt.Println("\xe4\xb8\x96") //输出“世”，因为底层\xe4\xb8\x96 对应的字就是世字

	fmt.Printf("############################3\n")
	fmt.Printf("%#v\n", []byte("hello, 世界")) //%#v打印相应值的Go语法表示，这里是打印的底层的数据存储形式[]byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c}

	fmt.Printf("############################4\n")
	//将字符串强行转换成byte字节数组就能进行遍历，但是bstr[i]表示的还是
	str := "hello, 世界"
	bstr := []byte(str)
	for i := 0; i < len(s); i++ {
		fmt.Printf("bstr[i]:%v\n", bstr[i])
	}

	for i, v := range bstr {
		fmt.Println(i, v)
	}

}
