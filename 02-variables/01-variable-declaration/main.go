package main

import "fmt"

func main() {
	// Declaring Variables
	var myStr string = "Hello"
	var myInt int = 100
	var myFloat float64 = 45.12
	fmt.Println(myStr, myInt, myFloat)


	//================================


	// Multiple Declarations
	var (
		employeeId int = 5
		firstName, lastName string = "Satoshi", "Nakamoto"
	)
	fmt.Println(employeeId, firstName, lastName)


	//================================


	// Short variable declaration syntax
	name := "Rajeev Singh"
	// 这行代码实际上做了三件事：
	// - 声明一个新变量 name
	// - 推断其类型为 string(根据右侧值)
	// - 赋值 "Rajeev Singh" 给这个变量
	age, salary, isProgrammer := 35, 50000.0, true

	fmt.Println(name, age, salary, isProgrammer)
}