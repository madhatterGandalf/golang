package main

import(
	"fmt"
)
//常量值必须是编译期可确定的数字、字符串、布尔值。

const a,b int = 12,34		//多常量的定义
const name  = "tom"		//常量的类型推断

//常量组
const(
	c,d = "abc","def"
	e bool = false
)

const (
	g int = 69
	h						//常量组中未赋值的 常量与上一个常量相同
)

//枚举
const (
	Sunday 		= iota 		//关键字 iota 定义常量组中从 0 开始按⾏计数的⾃增枚举值。
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

//利用iota的自增特性
const(
	_ 		  = iota
	KB  int64 = 1 << (10 * iota)
	MB
	GB
	TB
)

//同⼀常量组中，可以提供多个 iota，它们各自增长
const (
	i,j,k = iota,iota,iota			//iota 定义常量组中从 0开始
	l,m,n
)

//iota ⾃增被打断，须显式恢复
const(
	o 	= 	iota
	p
	q 	= 	45 						//打断
	r								//与q的值相同
	s 	= 	iota					//显式恢复。注意计数包含了 q、 r 两⾏
	t
)



func main(){
	const f = "ghi" 				//未使用的局部常量不会报错

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(g)
	fmt.Println(h)

	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)

	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)
	fmt.Println(t)

}
