package main

import "fmt"

func main() {
	/*
		数值类型（包括complex64/128）为 0

		布尔类型为 false

		字符串为 ""（空字符串）

		以下几种类型为 nil：
	*/

	// 第一种，指定变量类型，如果没有初始化，则变量默认为零值。
	var a int
	var b float32
	var c float64
	var d string
	var e bool
	fmt.Println(a, b, c, d, e)
	fmt.Println(d)

	var f *int
	var g []int
	var h map[string]int
	var i chan int
	var j func(string) int
	var k error //error是接口？
	fmt.Println(f, g, h, i, j, k)
}
