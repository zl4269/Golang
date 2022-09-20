package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("please visit http://127.0.0.1:12345")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := fmt.Sprintf("你好，世界！--Time:%s", time.Now().String())
		fmt.Fprintf(w, "%v\n", s) //区别:Printf(),Sprintf(),Fprintf()函数；
		log.Printf("%v\n", s)
	})
	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServe: ", err) //func log.Fatal(v ...interface{}),"ListenAndServe:",err类似于拼接

	}
}

/* 区别:Printf(),Sprintf(),Fprintf()函数；  //记：S表示字符串,F表示文件设备
① Printf() 是把格式化字符串输出到标准到标准输出（一般是屏幕，可以重定向）

②Sprintf() 是把格式化字符串输出到指定的字符串中，可以用一个变量来接受，然后在打印

③Fprintf() 是把格式字符串输出到指定的文件设备中，所以参数比Printf 多一个文件指针*File

主要用于文件操作，Fprintf() 是格式化输出到一个 Stream ,通常是一个文件 */
