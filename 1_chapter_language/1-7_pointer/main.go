package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var p int = 10
	var p1 *int							//⽀持指针类型 *T
	var p2 **int						//指针的指针 **T
	p1 = &p								//操作符 "&" 取变量地址
	p2 = &p1

	fmt.Printf("%v\n",p)
	fmt.Printf("%v\n",*p1)		//"*" 透过指针访问目标对象
	fmt.Printf("%v\n",**p2)


	//不能对指针做加减法等运算。
	var p3 int = 1
	var p4 *int
	p4 = &p3
	//p4++
	(*p4)++
	fmt.Println(p4)


	// 直接⽤指针访问目标对象成员，⽆须转换。
	type data struct{
		a int
	}
	var d = data{1234}
	var p5 *data
	p5 = &d
	fmt.Printf("%p, %v\n", p5, p5.a)


	//可以在 unsafe.Pointer 和任意类型指针间进⾏转换
	var p6 int = 3454563575
	p7 := unsafe.Pointer(&p6)
	n  := (*[3]byte)(p7)
	for i := 0;i<len(n);i++{
		fmt.Println("---",n[i])
	}


	//将 Pointer 转换成 uintptr，可变相实现指针运算
	d1 := struct {
		s string
		x int
	}{"abc", 100}

	p11 := uintptr(unsafe.Pointer(&d1)) 	// *struct -> Pointer -> uintptr
	p11 += unsafe.Offsetof(d1.x) 			// uintptr + offset
	p22 := unsafe.Pointer(p11) 				// uintptr -> Pointer
	px := (*int)(p22) 						// Pointer -> *int
	*px = 200 								// d.x = 200
	fmt.Printf("%#v\n", d)
	//注意： GC 把 uintptr 当成普通整数对象，它⽆法阻⽌ "关联" 对象被回收。


}
