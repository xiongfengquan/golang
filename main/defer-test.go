package main

import (
	"fmt"
)
/*
	defer 在一个函数中的个数不限，正常语句执行完再执行它们。若有多个defer,先顺序放到栈里，函数return前再从栈里面逐个取出执行（一句话总结先遇到的后执行）
	下面这段代码的输出为：

	start to echonum
	4
	3
	2
	1
	0
	end echonum
	c
	b
	a

*/
func echonum()  {
	for i := 0; i < 5; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func echotest(i string)  {
	fmt.Println(i)
}

func main()  {
	defer fmt.Println("a")
	defer echotest("b")
	defer fmt.Println("c")
	fmt.Println("start to echonum")
	echonum()
	fmt.Println("end echonum")
}
