package main

import (
	"bufio"
	"fmt"
	"os"
)
/*
	输出如下：

	请输入姓名：

	能不能好好写名字？
	xfq
	你好棒
	wwl
	你好美
	lxl
	你太抽象了，无法辨认
	go home
	bye


*/
func main()  {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入姓名：")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("系统繁忙，请重试。。。")
			continue
		}
		input = input[:len(input)-1]
		switch input {
		case "":
			fmt.Println("能不能好好写名字？")
		case "xfq":
			fmt.Println("你好棒")
		case "wwl":
			fmt.Println("你好美")
		case "go home":
			fmt.Println("bye")
			os.Exit(0)
		default:
			fmt.Println("你太抽象了，无法辨认")
		}
	}
}
