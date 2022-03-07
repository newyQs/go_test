package main

import "fmt"

func main() {
	/*
		常量是一个简单值的标识符，在程序运行时，不会被修改的量。
		常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
		常量的定义格式：
		const identifier [type] = value

		显示类型定义：const b string = "abc"
		隐示类型定义：const b = "abc"

		如果多个变量是同种类型：
		const c_name1, c_name2, c_name3
	*/
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Println("面积为：", area)
	fmt.Println(a, b, c)
}
