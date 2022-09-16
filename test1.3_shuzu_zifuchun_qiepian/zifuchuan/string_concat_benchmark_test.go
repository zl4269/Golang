//用于测试字符串拼接效率，方法：
//1.使用操作符 + ；  2.使用fmt.Sprintf;  3.使用strings.Join;  4.使用strings.Builder;  5.使用bytes.Buffer
package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var s1 []string = []string{
	"Rob Pike",
	"Robert Griesemer",
	"Ken Thompson",
}

func concatStringByOperator(s1 []string) string {
	var s string
	for _, v := range s1 {
		s += v
	}
	return s
}

func concatStringBySprintf(s1 []string) string { //效果最差
	var s string
	for _, v := range s1 {
		s = fmt.Sprintf("%s%s", s, v) //将v写到s上面，注意该函数的使用
	}
	return s
}

func concatStringByJoin(s1 []string) string { //用于字符串数组的拼接
	return strings.Join(s1, "")
}

func concatStringByStringBuilder(s1 []string) string {
	var b strings.Builder
	for _, v := range s1 {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByStringBuilderWithInitSize(s1 []string) string { //做了初始化，是所有方法中效率最高的
	var b strings.Builder
	b.Grow(64) //预估字符串的长度
	for _, v := range s1 {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByBytesBuffer(s1 []string) string {
	var b bytes.Buffer
	for _, v := range s1 {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByBytesBufferWithInitSize(s1 []string) string {
	buf := make([]byte, 0, 64)
	b := bytes.NewBuffer(buf)
	for _, v := range s1 {
		b.WriteString(v)
	}
	return b.String()
}

func BenchmarkConcatStringByOperator(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByOperator(s1)
	}
}

func BenchmarkConcatStringBySprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringBySprintf(s1)
	}
}

func BenchmarkConcatStringByJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByJoin(s1)
	}
}

func BenchmarkConcatStringByStringBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByStringBuilder(s1)
	}
}

func BenchmarkConcatStringByStringBuilderWithInitSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByStringBuilderWithInitSize(s1)
	}
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByBytesBuffer(s1)
	}
}

func BenchmarkConcatStringByBytesBufferWithInitSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByBytesBufferWithInitSize(s1)
	}    
}

/*
结果：
BenchmarkConcatStringByOperator-2                       10960780               107.5 ns/op            72 B/op          2 allocs/op
BenchmarkConcatStringBySprintf-2                         2169717               547.2 ns/op           160 B/op          8 allocs/op
BenchmarkConcatStringByJoin-2                           20500099                59.56 ns/op           48 B/op          1 allocs/op
BenchmarkConcatStringByStringBuilder-2                  10023708               119.9 ns/op            80 B/op          3 allocs/op
BenchmarkConcatStringByStringBuilderWithInitSize-2      23217510                51.20 ns/op           64 B/op          1 allocs/op
BenchmarkConcatStringByBytesBuffer-2                    13604835                88.68 ns/op          112 B/op          2 allocs/op
BenchmarkConcatStringByBytesBufferWithInitSize-2        20012224                58.42 ns/op           48 B/op          1 allocs/op
*/
