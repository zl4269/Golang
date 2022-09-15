//Go1.9版本之后引入了支持并非安全的sync.Map类型,使得map支持并发

package main

import (
	"fmt"
	"time"
)

func doIteration(m map[int]int) {
	for k, v := range m {
		s := fmt.Sprintf("[%d,%d]", k, v)
		fmt.Printf("s: %v\n", s)
	}

}

func doWrite(m map[int]int) {
	for k, v := range m {
		m[k] = v + 1
	}
}

func main() {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}

	go func() {
		for i := 0; i < 1000; i++ {
			doIteration(m)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			doWrite(m)
		}
	}()
	time.Sleep(5 * time.Second)

}
