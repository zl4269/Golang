//体验函数作为一等公民的使用部分方法
package main

import "fmt"

//函数作为一种数据类型(包含了参数和返回值类型)
type binaryCalcFunc func(int, int) int

func main() {
	//函数可以跟其他数据类型一样赋值给interface接口
	var i interface{} = binaryCalcFunc(func(x, y int) int { return x + y }) //将后面的内部函数转换成binaryCalcFunc类型，在赋值给interface接口
	v, ok := i.(binaryCalcFunc)                                             //进行类型断言
	if !ok {
		fmt.Println("type assertion error")
		return
	}
	fmt.Println(v(15, 5)) //将函数类型执行（传入参数）
	fmt.Println("############################# ")

	//将函数类型作为channel的类型
	//c := make(chan func(int,int) int,10)
	c := make(chan binaryCalcFunc, 10)
	c <- func(i1, i2 int) int { return i1 + i2 } //将函数写入channel
	f := <-c                                     //将函数写出，此时f的类型就是binaryCalcFunc
	fmt.Println(f(15, 5))
	fmt.Println("######################################")

	//将函数类型存储在slice中
	fns := []binaryCalcFunc{
		func(i1, i2 int) int { return i1 + i2 }, //注意i1和i2的作用域
		func(i1, i2 int) int { return i1 - i2 },
		func(i1, i2 int) int { return i1 * i2 },
		func(i1, i2 int) int { return i1 / i2 },
	}
	fmt.Println(fns[0](15, 5)) //切片中每个元素的值都是binaryCalcFunc
}
