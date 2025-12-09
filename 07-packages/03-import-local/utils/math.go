
package utils  // 注意：同一个包名

// Add 返回两个整数的和
func Add(a, b int) int {
    return a + b
}

// 可以调用同包的其他文件中的函数
func Multiply(a, b int) int {
    result := 0
    for i := 0; i < b; i++ {
        result = Add(result, a)  // 调用同包的Add函数
    }
    return result
}