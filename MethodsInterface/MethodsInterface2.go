//接口类型与结构体类型相关的嵌入式有三种组合：
//1)在接口类型中嵌入接口类型
//2)在结构体类型中嵌入接口类型
//3)在结构体类型汇总嵌入结构体类型
//注意：没有在接口类型中嵌入结构体类型

//这里介绍第一种：在接口类型中嵌入接口类型
//最经典的例子就是io包中的ReadWriter\ReadWriterClose等函数
package main

import (
	"MethodsInterface/utils"
	"io"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

//以上是三种基本接口类型，现在实现接口类型的一个嵌入
type ReadWriter interface {
	Reader //嵌人类型的时候只需要填入类型的名称
	Closer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

func main() {
	utils.DumpMethodSet((*io.Writer)(nil))
	utils.DumpMethodSet((*io.Reader)(nil))
	utils.DumpMethodSet((*io.Closer)(nil))
	utils.DumpMethodSet((*io.ReadWriter)(nil))
	utils.DumpMethodSet((*io.ReadWriteCloser)(nil))

}

/*
结果：
io.Writer`s method set:
- Write

io.Reader`s method set:
- Read

io.Closer`s method set:
- Close

io.ReadWriter`s method set:
- Read
- Write

io.ReadWriteCloser`s method set:
- Close
- Read
- Write
*/
