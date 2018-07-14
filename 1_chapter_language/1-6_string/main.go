package main

import "fmt"

func main(){
	var s1 string
	var s2 string = "abcd"
	//var s3 string = "efgh"

	fmt.Printf("string的默认值:%q\n",s1)
	fmt.Printf("string的数组访问:%v\n",s2[1])				//字符编码 98
	fmt.Printf("string的数组访问:%s\n",string(s2[1]))		//字符编码 98 转 string
	//fmt.Printf("string的数组访问取指针%q\n",&s3[1])


	//使⽤ "`" 定义不做转义处理的原始字符串，⽀持跨⾏。
	s := `a
		b\r\n\x00
		c`
	fmt.Println(s)

	s4 := "hello " +		//连接跨⾏字符串时， "+" 必须在上⼀⾏末尾，否则导致编译错误。
		"word"
	fmt.Println(s4)
	/*
	s5 := "Hello, "
	+ "World!" // Error: invalid operation: + untyped string
	*/


	//⽀持⽤两个索引号返回⼦串。⼦串依然指向原字节数组，仅修改了指针和⻓度属性。
	s5 := "Hello, World!"
	s6 := s5[:5] 		// Hello
	s7 := s5[7:] 		// World!
	s8 := s5[1:5] 		// ello

	fmt.Printf("%s---,%s----,%s",s6,s7,s8)


	//单引号字符常量
	fmt.Printf("%T\n", 'a')
	var c1, c2 rune = '\u6211', '们'						//rune 是 int32类型
	fmt.Println(c1,c2)
	fmt.Println(c1 == '我', string(c2) == "\xe4\xbb\xac")


	//修改字符串
	s9 := "abcd"
	bs := []byte(s9)			//转byte后会复制数据
	bs[1] = 'B'
	fmt.Println(string(bs))
	fmt.Printf("%p,---%p\n",&s9,bs)	//内存地址不一致

	u := "电脑"
	us := []rune(u)				//注意rune 转 byte 后 会无法表示 因为编码的数字超过byte可表示范围
	us[1] = '话'
	fmt.Println(string(us))


	//⽤ for 循环遍历字符串时，也有 byte 和 rune 两种⽅式。
	s10 := "abc汉字"
	for i := 0; i < len(s10); i++ { 		// byte
		fmt.Printf("%c,---%T\n", s10[i],s10[i])
	}
	fmt.Println()
	for _, r := range s10 { 				// rune
		fmt.Printf("%c,---%T\n", r,r)
	}
}
