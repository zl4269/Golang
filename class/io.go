//bufio包实现了io.Reader和io.Writer
//实现从终端读取数据

package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader

func main() {
	inputReader = bufio.NewReader(os.Stdin) //读取器（实现带缓冲的io），并返回指针。将读取器跟标准输出绑定  os是操作系统相关的库
	//func bufio.NewReader(rd io.Reader) *bufio.Reader  该函数中的实参可以是任何实现了io.Reader和io.Writer接口的参数
	fmt.Println("Please enter some input: ")
	input, err := inputReader.ReadString('\n') //func (*bufio.Reader).ReadString(delim byte) (string, error)
	/* 	方法从输入中读取内容，直到碰到 delim 指定的字符，然后将读取到的内容连同 delim 字符一起放到缓冲区。
	   	ReadString 返回读取到的字符串，如果碰到错误则返回 nil。如果它一直读到文件结束，则返回读取到的字符串和 io.EOF。
	   	如果读取过程中没有碰到 delim 字符，将返回错误 err != nil。 */   // \n表示回车，读取到回车就终止，然后返回err=nil 
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
		fmt.Printf("err: %v\n", err)
	}

}
