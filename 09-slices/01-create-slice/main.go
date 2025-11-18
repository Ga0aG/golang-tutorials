package main

import "fmt"

// 切片是一个 长度可变的数组。

func main() {
	// Creating a slice using a slice literal
	var s = []int{3, 5, 7, 9, 11, 13, 17} // Creates an array, and returns a slice reference to the array

	// Short hand declaration
	t := []int{2, 4, 8, 16, 32, 64}

	fmt.Println("s = ", s)
	fmt.Println("t = ", t)

	var slice1 []int = make([]int, 10)
	// load the array/slice:
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}
