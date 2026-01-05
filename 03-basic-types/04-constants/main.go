package main

import "fmt"

func main() {
	// Untyped Constant
	const myFavLanguage = "Python"
	const sunRisesInTheEast = true

	// Multiple declaration
	const country, code = "India", 91

	const (
		employeeId string  = "E101"
		salary     float64 = 50000.0
	)

	// Typed Constant
	const typedInt int = 1234
	const typedStr string = "Hi"

	fmt.Println(myFavLanguage, sunRisesInTheEast, country, code, employeeId, salary, typedInt, typedStr)


	// iota 是一个从 0 开始的整数常量生成器
  // 在每个 const 块中重置为 0
	const (
		ColorRed      = iota
		ColorOrange
		ColorYellow
	)

	const (
		Dog      = iota
		Cat
		Mouse
	)

	const (
		Monday      = iota + 1
		Tuesday
		Thursday = 4
		Happyday
		Friday   = iota + 1
		Saturday
	)
	fmt.Println(ColorYellow, Mouse) //2. 2
	fmt.Println(Monday, Tuesday, Thursday, Happyday, Friday, Saturday) // 1 2 4 4 5 6
}
