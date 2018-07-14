package main

import "fmt"

func main(){
	//可⽤ type 在全局或函数内定义新类型。
	type big int64
	var b big = 12
	fmt.Printf("%T\n",b)		//main.big


	type bigint int64
	x := 778
	const xx = 998
	var bi bigint = bigint(x)		//必须显示转换
	var bi2 bigint = xx				//常量不需要显示转换
	fmt.Printf("%T\n",bi)
	fmt.Printf("%T\n",bi2)

}
