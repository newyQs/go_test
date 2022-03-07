package main

import "fmt"

func main() {
	/*
		&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。	(A && B) 为 False
		||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。	(A || B) 为 True
		!	逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。	!(A && B) 为 True
	*/
	var a bool = true
	var b bool = false
	if (a && b) {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if (a || b) {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if (a && b) {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if (!(a && b)) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
}
