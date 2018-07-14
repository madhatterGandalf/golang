package main

import "fmt"

type Color int

//枚举
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

//使用ioat自增特性定义枚举
const (
	Black Color = iota
	Red
	Blue
)

func test(color Color){
	fmt.Println(color)
}

func main() {

	test(Black)

	test(1) 			// 常量会被编译器⾃动转换。

	i := 1					// 报错 i 不是常量值
	test(i)

	const i2 = 1
	test(i2)				//传入常量
}
