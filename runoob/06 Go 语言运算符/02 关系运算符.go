package main

import "fmt"

func main() {
	/*
		==	检查两个值是否相等，如果相等返回 True 否则返回 False。	(A == B) 为 False
		!=	检查两个值是否不相等，如果不相等返回 True 否则返回 False。	(A != B) 为 True
		>	检查左边值是否大于右边值，如果是返回 True 否则返回 False。	(A > B) 为 False
		<	检查左边值是否小于右边值，如果是返回 True 否则返回 False。	(A < B) 为 True
		>=	检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。	(A >= B) 为 False
		<=	检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。	(A <= B) 为 True
	*/
	var a int = 21
	var b int = 10

	if (a == b) {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}
	if (a < b) {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}

	if (a > b) {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if (a <= b) {
		fmt.Printf("第四行 - a 小于等于 b\n")
	}
	if (b >= a) {
		fmt.Printf("第五行 - b 大于等于 a\n")
	}

}
