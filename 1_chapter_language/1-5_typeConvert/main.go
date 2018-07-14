package main

import "fmt"

//1.5 类型转换


func main(){
	//向上转型
	var b byte = 100
	var c int = int(b)
	fmt.Printf("%T---,%T\n",b,c)

	//向下转型
	var bb int = 100
	var cc byte = byte(bb)
	fmt.Printf("%T---,%T\n",bb,cc)

	//类型转换
	var d  = map[int]string{11:"aa",22:"bb"}
	var e interface{}
	e = d

	f := e.(map[int]string)					//interface转map
	fmt.Printf("%T---,%T\n",e,f)


	//优先级问题
	/*
		*Point(p) 				// 相当于 *(Point(p))
		(*Point)(p)

		<-chan int(c) 			// 相当于 <-(chan int(c))
		(<-chan int)(c)
	*/

	var g = make(chan chan string,1)
	var h = make(chan string,1)
	h <- "dd"
	g <- h
	close(g)
	close(h)

	h <- "gg"		//close 后不能再次放入值

	// bool 类型

	/*
		a := 100
		if a { 	//不能将其他类型当 bool 值使⽤
			println("true")
		}
	*/
}