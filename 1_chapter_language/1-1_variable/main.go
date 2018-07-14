package main

import (
	"fmt"
)

//一次定义多个变量  当前定义方式可用于函数内部
var (
	e = 23
	f = "abc"
	g = 12.34
)

func main(){
	//常规定义变量的三种方式
	var a int					//不做初始化
	var b float32 = 1.235 		//定义时指定类型
	var c = 11					//自动类型识别
	fmt.Println(a) 				//0
	fmt.Println(b) 				//1.235
	fmt.Println(c) 				//11

	//函数内部的定义 (当前方式*不适用*于函数外部)
	d := 45.369 				//自动类型识别
	fmt.Println(d)				//45.369

	//测试多个变量的组合定义
	fmt.Println(e,f,g)

	//多变量赋值
	h,i := [3]int{0, 1, 2},"def"
	fmt.Println(h,i)			//先计算所有相关值，然后再从左到右依次赋值。

	//忽略值占位
	j,_ := placeholder(11,"ghi")			//编译器会将未使用的*局部变量*当做错误
	fmt.Println(j)

	//重新赋值与定义新同名变量的区别
	s := "abc"
	fmt.Println(&s)
	s, y := "hello", 20 							//重新赋值: 与前 s 在同⼀层次的代码块中，且有新的变量被定义。
	fmt.Println(&s, y) 								//通常函数多返回值 err 会被重复使⽤。
	{
		s, z := 1000, 30 							// 定义新同名变量: 不在同⼀层次代码块。
		fmt.Println(&s, z)
	}

}

func placeholder(as int,df string)(int,string){
	df = df + "jkl"
	as++
	return as,df
}