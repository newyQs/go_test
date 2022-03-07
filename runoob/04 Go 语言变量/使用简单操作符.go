package main

import "fmt"

var g int = 18

func main() {
	var a string = "hello-word"
	fmt.Print(a)

	b := "nihao"
	fmt.Print(b)

	//var c int = 12 // 函数体内的变量声明必须被使用，全局变量是可以被声明但不使用的
	fmt.Println("wohenhao ") //c declared but not used
}
