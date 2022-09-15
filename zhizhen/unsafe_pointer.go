package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	// v1 := uint(12)
	// v2 := int(13)

	// fmt.Println(reflect.TypeOf(v1)) //uint   //注意包反射 包反射实现了运行时反射，允许程序操作具有任意类型的对象。典型的用法是使用带有静态类型接口{}的值，并通过调用TypeOf提取其动态类型信息，TypeOf将返回type。
	// fmt.Println(reflect.TypeOf(v2)) //int

	// fmt.Println(reflect.TypeOf(&v1)) //*uint
	// fmt.Println(reflect.TypeOf(&v2)) //*int

	// p := &v1
	// p = (*uint)(unsafe.Pointer(&v2)) //使用unsafe.Pointer进行类型的转换

	// fmt.Println(reflect.TypeOf(p)) // *unit
	// fmt.Println(*p)                //13

	v1 := uint(11) //显示的初始化，不使用默认的数据类型
	v2 := int(12)

	//输出v1v2的类型和指针类型
	fmt.Printf("%v\n", reflect.TypeOf(v1))  //uint   //注意包反射 包反射实现了运行时反射，允许程序操作具有任意类型的对象。典型的用法是使用带有静态类型接口{}的值，并通过调用TypeOf提取其动态类型信息，TypeOf将返回type。
	fmt.Printf("%v\n", reflect.TypeOf(v2))  //int
	fmt.Printf("%v\n", reflect.TypeOf(&v1)) //*uint
	fmt.Printf("%v\n", reflect.TypeOf(&v2)) //*int

	//通过 unsafe.Pointer作为桥梁进行转换
	p := &v1
	p = (*uint)(unsafe.Pointer(&v2)) //将v2指针转换成uint
	fmt.Printf("typeof p: %v\n", reflect.TypeOf(p))
	fmt.Printf("value of p:%v\n", *p)
}
