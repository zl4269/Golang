//Go语言可以直接调用C代码。也可以引用C语言的外部库
package main

/*
#include<stdio.h>

int sayhelloc(int num){
	printf("hello C");
	return num;
}
*/
import "C" //注意一定要紧跟着C源码。不能留空格或者空行，否者会报错
import "fmt"

func sayhellogo() {
	fmt.Println("hello Go")
	fmt.Println(C.sayhelloc(3))
}

func main() {
	ans := C.sayhelloc(4)
	fmt.Println(ans)
	sayhellogo()
}
