package main

import (
	"fmt"
)

func getType(a interface{}) {
	switch a.(type) { //断言语法:x.(T) 其中x表示要断言的接口(默认数据类型也是interface)；T表示要判断的类型。特殊点：这里的type不用先定义，因为其要的断言的类型是变化的。
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float64")
	default:
		fmt.Println("unknow type")
	}
}

func main() {
	var a string
	a = "10"
	getType(a)
}
