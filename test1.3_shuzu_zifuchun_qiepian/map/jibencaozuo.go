//主要介绍map的基本操作
package main

import "fmt"

func main() {
	/* 	//1.map"零值不可用",零值为nil
	   	var m map[string]int
	   	m["key"] = 1
	   	fmt.Printf("m: %v\n", m)  //报错：ssignment to entry in nil map */

	//2.插入数据，map会进行内存管理，因此除非是系统内存耗尽，不然就不会报错
	//（创建map变量的两种方法：复合字面值和make创建）
	m := make(map[string]int) //区别零值，零值只是声明不初始化。这里用make创建了一个m的map变量
	m["key1"] = 1
	m["key2"] = 2
	fmt.Printf("m: %v\n", m) //m: map[key1:1 key2:2]
	m1 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Printf("m1: %v\n", m1) //m1: map[key1:1 key2:2]

	//3.统计元素个数：len()函数
	l := len(m)
	fmt.Printf("l: %v\n", l) //l: 2

	//4.判断某个元素是否在map中:使用 commma ok模式
	v, ok := m["key1"] //判断key1是否在m中
	if ok {
		fmt.Printf("v: %v\n", v) //v: 1
	}
	//一般不会直接读取，因为在map中，查找一个元素不存在，直接读取也会赋值一个0,所以判断一个元素是否存在于map最好使用comma ok模式
	vn := m["key3"]            //key3不存在也会赋值0
	fmt.Printf("vn: %v\n", vn) //vn: 0

	//5.删除数据，使用内置的delete()删除数据。如果数据不存在就回panic，所以删除之前要判断数据是否存在
	_, ok1 := m["key1"]
	if ok1 {
		delete(m, "key1")
	}
	fmt.Printf("m: %v\n", m)

}
