//这里介绍使用可变长参数实现选项功能：对于后期结构体的成员变量的增加，代码的的扩容简单
package main

import "fmt"

type FinishedHouse struct {
	style                 int    //0:chinese  1:American  2:Europe
	centraAirConditioning bool   //true 或 false
	floorMaterial         string //"ground-tile"或"wood"
	wallMaterial          string //"latex"或"paper"或"diatom-mud"
}

//定义选项实现的函数(作为返回值)
type Option func(*FinishedHouse)

//配置默认参数或者使用参数实例化FinishedHouse
func NewFinishedHouse(options ...Option) *FinishedHouse {
	//可以将*FinishedHouse作为一个参数传入，就能达到全局变量的效果
	//配置默认值
	h := &FinishedHouse{
		style:                 0,
		centraAirConditioning: true,
		floorMaterial:         "wood",
		wallMaterial:          "paper",
	}
	for _, option := range options {
		option(h)
	}
	return h
}

func WithStyle(style int) Option {
	return func(h *FinishedHouse) {
		h.style = style
	}
}

func WithWallMaterial(material string) Option {
	return func(h *FinishedHouse) {
		h.wallMaterial = material
	}
}

func WithFloorMaterial(material string) Option {
	return func(h *FinishedHouse) {
		h.floorMaterial = material
	}
}

func WithCentralAirConditioning(centralAirConditiong bool) Option {
	return func(h *FinishedHouse) {
		h.centraAirConditioning = centralAirConditiong
	}
}

func main() {
	fmt.Printf("%+v\n", NewFinishedHouse()) //使用默认值
	fmt.Printf("%+v\n", NewFinishedHouse(WithStyle(1), WithFloorMaterial("ground-tile"), WithCentralAirConditioning(false)))
}
