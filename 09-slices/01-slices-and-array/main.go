package main

import "fmt"

func main() {
	// Creation
	var a = [5]int{2, 4, 6, 8, 10} // 固定长数组
	var s = []int{3, 5, 7, 9, 11, 13, 17} // 变长数组
	s = append(s, 10)
	fmt.Println("a = ", a)
	fmt.Println("s = ", s)

	// Iteration
	aSum := 0
	for i := 0; i < len(a); i++ {
		aSum = aSum + a[i]
	}
	sSum := 0
	for i := 0; i < len(s); i++ {
		sSum = sSum + s[i]
	}
	fmt.Println("Sum of a = ", aSum)
	fmt.Println("Sum of s = ", sSum)

	// 切片引用数组
	s1 := a[:]
	a[0] = 10
	fmt.Println("s = ", s1) // s =  [10 4 6 8 10]
}