//1.如果某个自定义类型T的方法集合是某个接口类型的方法集合的超集，那么就说类型T实现了该接口，并且类型T变量可以赋值给该接口类型的变量。这就是方法集合决定接口的实现。
//要判断自定义类型是否实现了某个接口类型，就要识别自定义类型和接口类型的方法集合。

//这里探究:对于非接口类型的自定义类型T，其方法集合为所有receiver为T类型的方法组成
//         而类型*T的方法集合则包含所有receiver为T和*T类型的方法。

package main

import "MethodsInterface/utils" //当前目录路径开始就行

type Interface interface {
	M1()
	M2()
}

type T struct{}

func (t T) M1() {}

func (t *T) M2() {}

func main() {
	var t T                 //对于非接口类型的自定义类型T，其方法集合为所有receiver为T类型的方法组成
	var tp *T               //而类型*T的方法集合则包含所有receiver为T和*T类型的方法
	utils.DumpMethodSet(&t) //参数是地址
	utils.DumpMethodSet(&tp)
	utils.DumpMethodSet((*Interface)(nil)) //将nil转换成*Interface
}

/*
结果:
main.T`s method set:
- M1

*main.T`s method set:
- M1
- M2

main.Interface`s method set:
- M1
- M2
*/
