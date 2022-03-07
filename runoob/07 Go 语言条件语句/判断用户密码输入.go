package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Printf("请输入密码：   \n")
	fmt.Scan(&a)
	if a == 5211314 {
		fmt.Printf("请再次输入密码：")
		fmt.Scan(&b)
		if b == 5211314 {
			fmt.Printf("密码正确，门锁已打开")
		} else {
			fmt.Printf("非法入侵，已自动报警")
		}
	} else {
		fmt.Printf("非法入侵，已自动报警")
	}
}
