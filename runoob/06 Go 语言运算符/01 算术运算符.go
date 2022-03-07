package main

import "fmt"

func main() {
	/*
		+	相加	A + B 输出结果 30
		-	相减	A - B 输出结果 -10
		*	相乘	A * B 输出结果 200
		/	相除	B / A 输出结果 2
		%	求余	B % A 输出结果 0
		++	自增	A++ 输出结果 11
		--	自减	A-- 输出结果 9
	*/
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

}
