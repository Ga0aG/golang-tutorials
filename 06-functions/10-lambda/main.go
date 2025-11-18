package main

import "fmt"

func main() {
	f()
}
// 匿名函数同样被称之为闭包（函数式语言的术语）：它们被允许调用定义在其它环境下的变量
func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}