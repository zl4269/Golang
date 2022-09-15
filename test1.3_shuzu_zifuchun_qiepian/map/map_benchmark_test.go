//测试使用cap和不使用cap创建map的性能
package main

import (
	"math/rand"
	"testing"
	"time"
)

const mapSize = 10000

func generateWithCap() {
	rand.Seed(time.Now().UnixNano())
	m := make(map[int]int, mapSize)
	for i := 0; i < mapSize; i++ {
		m[i] = rand.Int()
	}
}

func generate() {
	rand.Seed(time.Now().UnixNano())
	m := make(map[int]int)
	for i := 0; i < mapSize; i++ {
		m[i] = rand.Int()
	}
}

func BenchmarkGenerateWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateWithCap()
	}
}

func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate()
	}
}

//运行时候的命令：go test -benchmem -bench=. map_benchmark_test.go
//可以看出，使用cap的效率大概是没有使用cap效率的两倍
