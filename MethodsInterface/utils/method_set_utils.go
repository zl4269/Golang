//这里实现了一个工具，用于输出自定义类型或接口类型的方法集合
package utils

import (
	"fmt"
	"reflect"
)

func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elemTyp := v.Elem()

	n := elemTyp.NumMethod()
	if n == 0 {
		fmt.Println("%s`s method set is empty!\n", elemTyp)
		return
	}

	fmt.Printf("%s`s method set:\n", elemTyp)
	for i := 0; i < n; i++ {
		fmt.Println("-", elemTyp.Method(i).Name)
	}
	fmt.Printf("\n")
}
