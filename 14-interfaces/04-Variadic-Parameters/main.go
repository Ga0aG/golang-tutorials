package main

import "fmt"

// 只能接收 int 类型
func sumInts(nums ...int) int {
	total := 0
	for _, n := range nums {
			total += n
	}
	return total
}

// 可以接收任意类型
// interface{} 是 Go 中的空接口，可以存储任何类型的值。在可变参数中使用 ...interface{} 是为了让函数能够接收任意类型、任意数量的参数。
func printAnything(values ...interface{}) {
	for _, v := range values {
			fmt.Printf("%v (type: %T)\n", v, v)
	}
}

func main() {
	// sumInts 只能接收 int
	sumInts(1, 2, 3)           // ✅ 正确
	// sumInts(1, "hello", 3)   // ❌ 编译错误

	// printAnything 可以接收任意类型
	printAnything(1, "hello", 3.14, true, []int{1, 2, 3})  // ✅ 正确
	// 输出：
	// 1 (type: int)
	// hello (type: string)
	// 3.14 (type: float64)
	// true (type: bool)
	// [1 2 3] (type: []int)
}