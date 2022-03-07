package main

import "fmt"

var x, y int
var (
	a int
	b bool
)
var c, d int = 1, 2
var e, f = 123, "hello"

func main() {
	// 注意，这种不带声明格式的只能出现在函数体里面
	g, h := 123, "hello"
	fmt.Print(x, y, a, b, c, d, e, f, g, h)
}
