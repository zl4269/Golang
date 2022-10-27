//这里介绍实现选项模式的第一种方法：通过参数暴露配置选项
//缺点就是：该接口无法实现扩容(写死了)。如果需要扩容，就要重写API
package main

import "fmt"

type FinishedHouse struct {
	style                 int    //0:chinese  1:American  2:Europe
	centraAirConditioning bool   //true 或 false
	floorMaterial         string //"ground-tile"或"wood"
	wallMaterial          string //"latex"或"paper"或"diatom-mud"
}

func NewFinishedHouse(style int, centralAirConditioning bool, floorMaterial, wallMater string) *FinishedHouse {
	//没有设置默认参数
	h := &FinishedHouse{
		style:                 style,
		centraAirConditioning: centralAirConditioning,
		floorMaterial:         floorMaterial,
		wallMaterial:          wallMater,
	}
	return h
}

func main() {
	fmt.Printf("%+v\n", NewFinishedHouse(0, true, "wood", "paper"))
}
