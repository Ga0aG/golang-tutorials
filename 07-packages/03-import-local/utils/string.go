
package utils  // 注意：同一个包名

// Reverse 反转字符串
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
// 调用同包math.go中的函数
func CalculateStringLength(s string) int {
    // 这里演示同包文件间的调用
    // 注意：同一个包下直接调用，不需要import
    return len(s)
}