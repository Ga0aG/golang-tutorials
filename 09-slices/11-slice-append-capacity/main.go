package main

import "fmt"

func main() {
	// Appending to a slice that has enough capacity to accommodate new elements
	slice1 := make([]string, 3, 5)
	copy(slice1, []string{"C", "C++", "Java"})

	// 当使用 append() 向切片添加元素且容量不足时，Go 会自动扩容。
	slice2 := append(slice1, "Python", "Ruby", "Go")

	fmt.Printf("slice1 = %v, len = %d, cap = %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 = %v, len = %d, cap = %d\n", slice2, len(slice2), cap(slice2))

	slice1[0] = "C#"
	fmt.Println("\nslice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	/*
		In this case, no new array would be allocated, and the elements would be added to the same underlying array.
	*/
}
