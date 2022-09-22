// receiver的一个经典难以理解的例子
package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() { //这里是取的地址
	fmt.Println(p.name) //无论是T类型实例还是*T类型实例，都既可以调用receiver为T类型的方法，也可以调用receiver为*T类型的方法。
}

func main() {
	//一下两端代码一段一段运行，就能达到效果
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 { //v定义在for循环的外面，由于采用的是*T，每次都是将切片的地址直接赋值给v，因为切片中的元素就是*T
		go v.print()
	}
	//结果(可能有所差异)
	/*
		three
		two
		one
	*/

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go v.print() //v相当于定义在for循环外面。由于采用的是*T,for循环每次取的是v的地址，再将切片中的值传入的，而不是取的切片的地址
	}
	/*
	   结果:
	   	six
	   six
	   six
	*/

	time.Sleep(time.Second * 5) //让主协程睡眠，使得goruntine运行
}
