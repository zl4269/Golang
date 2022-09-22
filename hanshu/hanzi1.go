package main

import "fmt"

//定义函子必需实现一种方法
type Slicbyte interface {
	Fmap(f func(byte) byte) Slicbyte
}

//定义函子的载体
type slicebyte struct {
	bytes []byte
}

//函子载体去实现方法
func (s slicebyte) Fmap(f func(byte) byte) Slicbyte {
	news := make([]byte, len(s.bytes))
	for i, v := range s.bytes {
		res := f(v)
		news[i] = res
	}
	return slicebyte{bytes: news}
}

//用数组对函子载体进行初始化
func newInit(s []byte) slicebyte {
	return slicebyte{bytes: s}
}

func main() {
	b := []byte("hello")
	fmt.Printf("b: %v\n", b)

	f := newInit(b)
	fmt.Printf("f: %v\n", f)

	fangfa := func(b byte) byte {
		return byte(b + 1)
	}

	newf := f.Fmap(fangfa) //fangfa调用的时候自己会传参
	fmt.Printf("newf:%v\n", newf)
}
