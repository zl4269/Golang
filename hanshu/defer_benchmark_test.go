package main

import "testing"

func sum(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += i
	}
	return total
}

func fooWithDefer() {
	defer func() {
		sum(10)
	}()
}

func fooWithoutDefer() {
	sum(10)
}

func BenchmarkFooWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithDefer()
	}
}

func BenchmarkFooWithoutDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithoutDefer()
	}
}

//执行：go test -bench . defer_benchmark_test.go
/*
结果：
goos: linux
goarch: amd64
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
BenchmarkFooWithDefer-2         149263591                8.022 ns/op
BenchmarkFooWithoutDefer-2      207186447                5.104 ns/op
发现：使用defer会使得性能下降一些。
*/
