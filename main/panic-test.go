package main

import (
	"errors"
	"fmt"
)

/*
	panic 引起的恐慌会随着调用链层层上抛，直至整个程序崩溃
	这里需要注意一点，即使引发了恐慌，恐慌前的defer也会执行完成，恐慌后的不再执行
	以下是返回：

	0
	3
	test defer
	panic: panic test

	goroutine 1 [running]:
	main.testPanic()
			/home/xf/docker-env/code/golang/src/golang/main/panic-test.go:16 +0x144
	main.main()
			/home/xf/docker-env/code/golang/src/golang/main/panic-test.go:21 +0x7e
	exit status 2

*/

func testPanic()  {
	defer fmt.Println("test defer")  //会执行
	fmt.Println("3")
	panic(errors.New("panic test"))
	fmt.Println("1")   //不会执行
}
func main()  {
	fmt.Println("0")
	testPanic()
	fmt.Println("2")  //不会执行
}
