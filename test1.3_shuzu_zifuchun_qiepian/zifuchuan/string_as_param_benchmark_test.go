//用于测试字符串传参时传入文件描述符和传入字符串指针效果是一样的
package main

import "testing"

var s string = `abcdef
ghijklmn
opqrst
uvwxyz`

func handleString(s string) {
	_ = s + "hello world"
}

func handlePtrToString(s *string) {
	_ = *s + "hello world"
}

func BenchmarkHandleString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		handleString(s)
	}

}

func BenchmarkHandlePtrToString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		handlePtrToString(&s)
	}

}

//命令行：go test -bench . -benchmem string_as_param_benchmark_test.go
/* //结果
BenchmarkHandleString-2         23742303                48.16 ns/op           48 B/op          1 allocs/op
BenchmarkHandlePtrToString-2    25135844                48.69 ns/op           48 B/op          1 allocs/op
效果是一样的 */
