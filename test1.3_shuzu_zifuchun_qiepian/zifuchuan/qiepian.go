package main

import "fmt"

func main() {
	//字符串的切片
	s := "hello, world"
	hello := s[:5]
	world := s[7:]

	fmt.Printf("############################1\n")
	fmt.Printf("hello:%v\n", hello) //hello:hello
	fmt.Printf("world:%v\n", world) //world:world

	fmt.Printf("############################2\n")
	//上述中的s可以直接使用字符串表示
	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Printf("s1:%v\n", s1) //s1:hello
	fmt.Printf("s2:%v\n", s2) //s2:world

}
