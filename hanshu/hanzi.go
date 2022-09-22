// 函子的使用
/*
函子必需满足两个条件
1.函子本身就是一个容器类型，如切片\map\channel
2.该容器类型需要实现一个方法，该方法接收一个函数类型的参数，并在容器中的每个元素上都会调用那个函数，得到一个新函子，原函子容器中元素与新函子容器的元素互不影响
*/
package main

import (
	"fmt"
)

type IntSliceFunctor interface {
	Fmap(fn func(int) int) IntSliceFunctor //定义函子实现的方法，该方法的参数是一个函数fn，注意函数作为参数的时候需要写参数和返回值的
}

type intSliceFunctorImpl struct { //函子的载体，初始化之后返回一个函子
	ints []int
}

func (isf intSliceFunctorImpl) Fmap(fn func(int) int) IntSliceFunctor {
	newInts := make([]int, len(isf.ints))
	for i, elt := range isf.ints {
		retInt := fn(elt)
		newInts[i] = retInt
	}
	return intSliceFunctorImpl{ints: newInts}
}

func NewIntSliceFunctor(slice []int) IntSliceFunctor {
	return intSliceFunctorImpl{ints: slice}
}

func main() {
	// 原切片
	intSlice := []int{1, 2, 3, 4}
	fmt.Printf("init a functor from int slice: %#v\n", intSlice)
	f := NewIntSliceFunctor(intSlice)
	fmt.Printf("original functor: %+v\n", f)

	mapperFunc1 := func(i int) int {
		return i + 10
	}

	mapped1 := f.Fmap(mapperFunc1)
	fmt.Printf("mapped functor1: %+v\n", mapped1)

	mapperFunc2 := func(i int) int {
		return i * 3
	}
	mapped2 := mapped1.Fmap(mapperFunc2)
	fmt.Printf("mapped functor2: %+v\n", mapped2)
	fmt.Printf("original functor: %+v\n", f) // 原functor没有改变
	fmt.Printf("composite functor: %+v\n", f.Fmap(mapperFunc1).Fmap(mapperFunc2))
}
