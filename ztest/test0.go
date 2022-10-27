package main

import "fmt"

func mymax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func mymin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func test(n, x int, a []int) int {
	var ans = 0
	var H = a[0] + x
	var L = a[0] - x
	for i := 0; i < n; i++ {
		var newH = a[i] + x
		var newL = a[i] - x
		if newH < L || newL > H {
			ans++
			L = newL
			H = newH
		} else {
			L = mymax(L, newL)
			H = mymin(H, newH)
		}

	}
	return ans
}

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	var a []int
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		a = append(a, tmp)
	}
	ans := test(n, x, a)
	fmt.Println(ans)

}
