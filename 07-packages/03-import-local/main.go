package main

import (
    "fmt"
    "myproject/utils"       // 1. import 本地别的package
    "myproject/config"      // 导入另一个本地包
)

/*
STEP1: go mod init myproject
STEP2: go run main.go
如果要编译的话
STEP3: go build -o build/app
STEP4: ./build/app
*/

func main() {
    fmt.Println("=== 使用 utils 包 ===")

    // 使用 utils 包中的函数
    sum := utils.Add(10, 20)
    fmt.Printf("Add(10, 20) = %d\n", sum)

    product := utils.Multiply(5, 6)
    fmt.Printf("Multiply(5, 6) = %d\n", product)

    // 使用 string 包中的函数
    reversed := utils.Reverse("Hello")
    fmt.Printf("Reverse('Hello') = %s\n", reversed)

    fmt.Println("\n=== 使用 config 包 ===")

    // 使用 config 包
    cfg := config.LoadConfig()
    fmt.Printf("Config: %+v\n", cfg)
    fmt.Printf("Database URL: %s\n", cfg.GetDatabaseURL())
}