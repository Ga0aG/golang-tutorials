package main

import (
	"fmt"
)

type Student struct {
	RollNumber int
	Name       string
}

func main() {
	// instance of student struct type
	s := Student{11, "Jack"}

	// Pointer to the student struct
	ps := &s
	fmt.Println(ps)

	// Accessing struct fields via pointer
	fmt.Println((*ps).Name)
	fmt.Println(ps.Name) // Same as above: No need to explicitly dereference the pointer

	// 在 Go 语言中这叫 选择器 (selector)。无论变量是一个结构体类型还是一个结构体类型指针，都使用同样的 选择器符 (selector-notation) 来引用结构体的字段：
	ps.RollNumber = 31
	fmt.Println(ps)
	s.RollNumber = 20
	fmt.Println(ps)
}
