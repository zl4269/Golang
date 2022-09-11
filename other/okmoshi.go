package main

import "fmt"

func main() {
	m := map[string]int{"key1": 1, "key2": 2}
	v, ok := m["key3"]
	if ok {
		fmt.Printf("v: %v\n", v) //注意返回值是map中的value
	} else {
		fmt.Println("key1 not found")
	}
}
