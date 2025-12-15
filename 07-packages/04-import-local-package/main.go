package main

import (
	"fmt"
	"myproject/utils" // 导入 03-import-local 项目中的 utils 包
)

// go mod tidy
// go run main.go

func main() {
	fmt.Println("=== 使用本地包 03-import-local 中的函数 ===")

	// 使用 utils 包中的 Add 函数
	sum := utils.Add(10, 20)
	fmt.Printf("10 + 20 = %d\n", sum)

	// 使用 utils 包中的 Multiply 函数
	product := utils.Multiply(5, 6)
	fmt.Printf("5 * 6 = %d\n", product)

	// 使用 utils 包中的 Reverse 函数
	text := "Hello, 世界"
	reversed := utils.Reverse(text)
	fmt.Printf("原字符串: %s\n", text)
	fmt.Printf("反转后: %s\n", reversed)

	// 使用 utils 包中的 CalculateStringLength 函数
	length := utils.CalculateStringLength(text)
	fmt.Printf("字符串长度: %d\n", length)
}
