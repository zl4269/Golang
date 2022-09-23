//接口类型与结构体类型相关的嵌入式有三种组合：
//1)在接口类型中嵌入接口类型
//2)在结构体类型中嵌入接口类型
//3)在结构体类型汇总嵌入结构体类型
//注意：没有在接口类型中嵌入结构体类型

//这里介绍第一种：在结构体中嵌入结构体
//注意两点：
//1)新的结构体能够调用被嵌入结构体的所有方法。
//2)虽然新的结构体能够调用所有的方法，但是对于新的结构体，T和T*类型的方法集合还是有差被的，取决于嵌入的形式
如：嵌入体{T1  *T2}
则：T类型的方法集合 = T1 的方法集合 + *T2的方法集合
   *T类型的方法集合 = *T1 的方法集合 + *T2的方法集合
package main

import (
	"MethodsInterface/utils"
	"fmt"
)

type T1 struct{}

func (T1) T1M1()   { println("T1`s M1") }
func (T1) T1M2()   { println("T1`s M2") }
func (*T1) PT1M3() { println("PT1`s M3") }

type T2 struct{}

func (T2) T2M1()   { println("T2`s M1") }
func (T2) T2M2()   { println("T2`s M2") }
func (*T2) PT2M3() { println("PT2`s M3") }

type T struct {
	T1
	*T2
}

func main() {
	t := T{
		T1: T1{}, //结构体中嵌入结构体的初始化
		T2: &T2{},
	}

	fmt.Println("call method through pt:")
	t.T1M1()
	t.T1M2()
	t.PT1M3()
	t.T2M1()
	t.T2M2()
	t.PT2M3()

	fmt.Println("\ncall method through pt:")
	pt := &t
	pt.T1M1()
	pt.T1M2()
	pt.PT1M3()
	pt.T2M1()
	pt.T2M2()
	pt.PT2M3()

	var t1 T1
	var pt1 *T1
	utils.DumpMethodSet(&t1)
	utils.DumpMethodSet(&pt1)

	var t2 T2
	var pt2 *T2
	utils.DumpMethodSet(&t2)
	utils.DumpMethodSet(&pt2)

	utils.DumpMethodSet(&t)
	utils.DumpMethodSet(&pt)
}

/*
结果：
call method through pt:
T1`s M1
T1`s M2
PT1`s M3
T2`s M1
T2`s M2
PT2`s M3

call method through pt:
T1`s M1
T1`s M2
PT1`s M3
T2`s M1
T2`s M2
PT2`s M3
main.T1`s method set:
- T1M1
- T1M2

*main.T1`s method set:
- PT1M3
- T1M1
- T1M2

main.T2`s method set:
- T2M1
- T2M2

*main.T2`s method set:
- PT2M3
- T2M1
- T2M2

main.T`s method set:
- PT2M3
- T1M1
- T1M2
- T2M1
- T2M2

*main.T`s method set:
- PT1M3
- PT2M3
- T1M1
- T1M2
- T2M1
- T2M2
*/
