package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func main() {
	//数组指针
	var a = [...]int{1, 2, 3}
	var b = &a //b就是数组a的指针，数组指针的复制只会复制一个，不像数组一样，数组指针的使用跟数组的使用是一样的
	fmt.Println(a[0], a[1])
	fmt.Println(b[1], b[2])

	fmt.Println("############################1")
	//遍历数组的方法：使用for range 效率最高，因为可以省略是否越界的判断---》一定是不会越界的
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Println("############################2")
	//函数数组
	//var decoder1 [2]func(io.Reader)(image.Image,error)
	var decoder2 = [...]func(io.Reader) (image.Image, error){ //{}里面跟着是初始化
		png.Decode,
		jpeg.Decode,
	}
	fmt.Printf("len of decoder2:%d\n", len(decoder2)) //结果为2

	fmt.Println("############################3")
	//接口数组
	//var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{ //接口interface
		//初始化
		123,
		"你好",
	}

	fmt.Printf("len of unknown2:%d\n", len(unknown2))

}
