package main

import "fmt"

func main() {
	//  var identifier type
	var a string = "runoob"
	fmt.Println(a)

	// 一次性声明多个
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 声明一个变量并初始化
	var d = "runood"
	fmt.Println(d)

	// 声明变量未初始化，默认为0
	var e int
	fmt.Println(e)

	// bool值未初始化，默认为0
	var f bool
	fmt.Println(f)

	var g float32
	fmt.Println(g)


}
