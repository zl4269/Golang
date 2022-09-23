//接口类型与结构体类型相关的嵌入式有三种组合：
//1)在接口类型中嵌入接口类型
//2)在结构体类型中嵌入接口类型
//3)在结构体类型汇总嵌入结构体类型
//注意：没有在接口类型中嵌入结构体类型

//这里介绍第一种：在结构体类型中嵌入接口类型

package main

import "MethodsInterface/utils"

type Interface interface {
	M1()
	M2()
}

type T struct {
	Interface //T类型中包含了接口类型
}

func (T) M3() {} // 没有用到 t T中的t，可以省略。

func main() {
	utils.DumpMethodSet((*Interface)(nil))

	var t T
	var tp *T
	utils.DumpMethodSet(&t)  //实现了嵌入的接口类型
	utils.DumpMethodSet(&tp) //*T会实现所有T和*T的方法

}

/*
main.Interface`s method set:
- M1
- M2

main.T`s method set:
- M1
- M2
- M3

*main.T`s method set:
- M1
- M2
- M3
*/
