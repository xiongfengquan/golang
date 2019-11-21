package main

import "fmt"

/*
	recover 可以拦截恐慌，重新获取程序控制权，让程序继续执行
	输出结果如下：

	0
	3
	test defer
	拦截恐慌，继续执行

*/

func panics()  {
	defer fmt.Println("test defer")  //会执行
	fmt.Println("3")
	panic("test panic")
	fmt.Println("4")  //以下的不会继续执行
}
func main()  {
	defer func() {
		if e:= recover(); e != nil {
			fmt.Println("拦截恐慌，继续执行")
		}
	}()
	fmt.Println("0")
	panics()
	fmt.Println("2")  //又可以继续执行了
}