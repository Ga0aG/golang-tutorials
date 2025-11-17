package main

import "fmt"

func main() {
	var (
		firstName, lastName string
		age                 int
		salary              float64
		isConfirmed         bool
	)

	// == 各种类型的零值一览表==
	// 类型	    零值	  说明
	// 数值类型	0	      包括整数和浮点数
	// 布尔类型	false
	// 字符串类型	"" (空字符串)
	// 指针类型	nil	    表示不指向任何内存地址
	// 切片类型	nil
	// 映射类型	nil
	// 通道类型	nil
	// 函数类型	nil
	// 接口类型	nil
	// 结构体类型	各字段的零值

	fmt.Printf("firstName: %s, lastName: %s, age: %d, salary: %f, isConfirmed: %t\n",
		firstName, lastName, age, salary, isConfirmed)
}
