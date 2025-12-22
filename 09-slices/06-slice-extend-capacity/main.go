package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original Slice")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	// cap() 可以测量切片最长可以达到多少：它等于切片的长度 + 数组除切片之外的长度
	s = s[1:5]
	fmt.Println("\nAfter slicing from index 1 to 5")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	// 多个切片如果表示同一个数组的片段，它们可以共享数据；因此一个切片和相关数组的其他切片是共享存储的
	s = s[:8]
	fmt.Println("\nAfter extending the length [:8]")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[2:]
	fmt.Println("\nAfter dropping the first two elements")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	var s1 = s[2:]
	fmt.Println("\nAfter dropping the first two elements")
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1))
}
