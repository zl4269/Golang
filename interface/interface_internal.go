//探究：空接口类型变量与非空接口类型变量之间的等值比较
package main

import "fmt"

type T interface {
}

func PrintEmptyInterfaceAndNonEmptyInterface() {
	var eif interface{} = T(5)
	var err = fmt.Errorf("5")

	println("eif:", eif)
	println("err:", err)
	println("eif=err:", eif == err)

	err = fmt.Errorf("6")
	println("eif:", eif)
	println("err:", err)
	println("eif=err:", eif == err)
}

func main() {
	PrintEmptyInterfaceAndNonEmptyInterface()
}
