//无论是T类型实例还是*T类型实例，都既可以调用receiver为T类型的方法，也可以调用receiver为*T类型的方法。
package main

import "fmt"

type T struct {
	a int
}

func (t T) M1() {
	t.a = 1
	fmt.Println(t.a)
}

func (t *T) M2() {
	t.a = 2
	fmt.Println(t.a)
}

func main() {
	var t T
	t.M1()
	t.M2() //T的实例t可以调用receiver为 *T的方法  相当于 (&t).M2()

	var tp = &T{} //注意指针的使用要显示赋零值，不赋值的话不会分配指定，导致地址无效
	tp.M1()       //*T的实例tp也可以调用receiver为T的方法  ，相当于(*tp).M1()  将tp解引用得到T的实例
	tp.M2()

}

/*
//结果：
1
2
1
2
*/
