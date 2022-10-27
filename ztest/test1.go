package main

import "fmt"

func main() {

	var slice []int
	var size int
	fmt.Scan(&size)
	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}
	fmt.Println(slice)
}
