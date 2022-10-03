//这里介绍实现选项模式的第一种方法：使用结构体封装配置选项（封装化原则）。将选项重现封装，使得传入参数是封装的结构体，保证了函数的传入参数是不变的
//缺点：传递nil和&Option{}的区别；每次调用Options中的字段时候也需要显示赋值，后继更改的代码也是很多
package main

import "fmt"

type FinishedHouse struct {
	style                 int    //0:chinese  1:American  2:Europe
	centraAirConditioning bool   //true 或 false
	floorMaterial         string //"ground-tile"或"wood"
	wallMaterial          string //"latex"或"paper"或"diatom-mud"
}

//重新封装结构体
type Options struct {
	Style                 int    //0:chinese  1:American  2:Europe
	CentraAirConditioning bool   //true 或 false
	FloorMaterial         string //"ground-tile"或"wood"
	WallMaterial          string //"latex"或"paper"或"diatom-mud"
}

func NewFinishedHouse(options *Options) *FinishedHouse {
	//如果默认值为nil，就会使用默认的参数
	var style int = 0
	var centraAirConditioning = true
	var floorMaterial = "wood"
	var wallMaterial = "paper"

	//重新封装之后，只要传入一个结构体就行
	if options != nil {
		style = options.Style
		centraAirConditioning = options.CentraAirConditioning
		floorMaterial = options.FloorMaterial
		wallMaterial = options.WallMaterial
	}

	h := &FinishedHouse{
		style:                 style,
		centraAirConditioning: centraAirConditioning,
		floorMaterial:         floorMaterial,
		wallMaterial:          wallMaterial,
	}

	return h
}

func main() {
	var p = &Options{}
	fmt.Printf("%+v\n", NewFinishedHouse(nil))
	fmt.Printf("%+v\n", NewFinishedHouse(p))
	fmt.Printf("%+v\n", NewFinishedHouse(&Options{Style: 1, CentraAirConditioning: false, FloorMaterial: "ground-tile", WallMaterial: "paper"}))
}

/*
结果：
&{style:0 centraAirConditioning:true floorMaterial:wood wallMaterial:paper}  注意:传入nil和传入&Options{}的区别
&{style:0 centraAirConditioning:false floorMaterial: wallMaterial:}
&{style:1 centraAirConditioning:false floorMaterial:ground-tile wallMaterial:pape
*/
