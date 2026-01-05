package main

import "fmt"

// 数组是具有相同 唯一类型 的一组已编号且**长度固定**的数据项序列（这是一种同构的数据结构）

func main() {
	// ===== Array Declaration =====
	fmt.Println("===== Array Declaration =====")
	var x [5]int // An array of 5 integers
	fmt.Println("x = ", x)

	var y [8]string // An array of 8 strings
	fmt.Println("y = ", y)

	var z [3]complex128 // An array of 3 complex numbers
	fmt.Println("z = ", z)
	// By default, all the array elements are assigned the zero value of the array type.
	// For example, if we declare an integer array, all the elements will be initialized with zero.
	// If we declare a string array, all the elements will be initialized with an empty string, and so on.

	// ===== Array Indexing =====
	fmt.Println("\n===== Array Indexing =====")
	var arr [5]int
	arr[0] = 100
	arr[1] = 101
	arr[3] = 103
	arr[4] = 105

	fmt.Printf("arr[0] = %d, arr[1] = %d, arr[2] = %d\n", arr[0], arr[1], arr[2])
	fmt.Println("arr = ", arr)

	// ===== Array Initialization =====
	fmt.Println("\n===== Array Initialization =====")
	// Declaring and initializing an array at the same time
	var a = [5]int{2, 4, 6, 8, 10}
	fmt.Println("a = ", a)

	// Short hand declaration for declaring and initializing an array
	b := [5]int{2, 4, 6, 8, 10}
	fmt.Println("b = ", b)

	// You don't need to initialize all the elements of the array.
	// The un-initialized elements will be assigned the zero value of the corresponding array type
	c := [5]int{2}
	fmt.Println("c = ", c)

	// Letting Go compiler infer the length of the array
	d := [...]int{3, 5, 7, 9, 11, 13, 17}
	fmt.Println("d = ", d)

	// ===== Array as Value Types =====
	fmt.Println("\n===== Array as Value Types =====")
	a1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
	a2 := a1 // A copy of the array `a1` is assigned to `a2`

	a2[1] = "German"

	fmt.Println("a1 = ", a1) // The array `a1` remains unchanged
	fmt.Println("a2 = ", a2)
}
