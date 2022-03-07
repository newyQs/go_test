package main

import "fmt"

func main() {
	// 第三种，如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误，格式：
	// v_name:=value
	var intVal int
	//intVal := 1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明
	fmt.Println(intVal)

	intVal2 := 2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	fmt.Println(intVal2)

	// 因此 intVal := 1 相当于 var intVal int \n intval = 1

	//可以将 var f string = "Runoob" 简写为 f := "Runoob"：
	f := "Runoob" // var f string = "Runoob"
	fmt.Println(f)
}
