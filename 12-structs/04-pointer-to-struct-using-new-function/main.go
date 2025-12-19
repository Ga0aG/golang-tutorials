package main

import "fmt"

type Employee struct {
	Id   int
	Name string
}

func main() {
	// You can also get a pointer to a struct using the built-in new() function
	// It allocates enough memory to fit a value of the given struct type, and returns a pointer to it
	pEmp := new(Employee)

	pEmp.Id = 1000
	pEmp.Name = "Sachin"

	fmt.Println(pEmp)

	// &struct 是一种简写，底层仍然会调用 new()，这里值的顺序必须按照字段顺序来写。
	pEmp2 := &Employee{100, "Sara"}
	fmt.Println(pEmp2)

	// Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体
}
