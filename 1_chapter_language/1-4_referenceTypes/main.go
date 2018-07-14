package main

import "fmt"

//slice、 map 和 channel  是引用类型
func referTypes(){
	sli := []int{1,2,3,4}
	ma  := map[int]string{12:"ss",13:"dd",14:"ff"}
	ch  := make(chan string)

	fmt.Printf("%p---,%p---,%p\n",sli,ma,ch)

	var sliCopy []int
	var maCopy	map[int]string
	var chCopy	chan string

	//执行copy
	sliCopy = sli		//引用赋值
	maCopy  = ma		//引用赋值
	chCopy  = ch		//引用赋值

	fmt.Printf("%p---,%p---,%p\n",sliCopy,maCopy,chCopy)

	//0xc042010280---,0xc042068060---,0xc04203c060
	//0xc042010280---,0xc042068060---,0xc04203c060

	//结论 copy后地址值完全一样
}

//make 和 new 的区别
func newAndMake(){
	a := [3]int{1,2,3}
	a[1] = 10

	b := make([]int,3)		//make 返回对象⽽⾮指针
	b[1] = 11

	//c := new([]int)		//new  返回指针
	//c[1] = 12
}

func main(){
	referTypes()

	newAndMake()
}
