//主要证明map是引用类型（作为函数参数的时候会改变map内部的值）
package main

import "fmt"

func foo(m map[string]int) {
	m["key1"] = 11
	m["key2"] = 12

}

func main() {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Printf("m: %v\n", m) //m: map[key1:1 key2:2]
	foo(m)
	fmt.Printf("m: %v\n", m) //m: map[key1:11 key2:12]  说明map作为函数参数的时候是引用传递

}
