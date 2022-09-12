package main

import "fmt"

func main() {
	var x interface{}
	x = 10 //interface能够隐式转换成任何接口或者类型
	/*上面不能可以直接写成：x := 10或者其他的定义形式，因为断言是断言interface接口值*/

	value, ok := x.(int)
	fmt.Printf("value:%d,ok:%v\n", value, ok)

	//断言的注意事项：最好返回第二个参数，如果不接收第二个参数也就是上面代码中的 ok，断言失败时会直接造成一个 panic。如果 x 为 nil 同样也会 panic。
	var str interface{} //str为nil,一定产生断言失败
	v := str.(string)
	fmt.Println(v)
}
