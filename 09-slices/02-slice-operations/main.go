package main

import "fmt"

// 切片操作：修改、长度容量、追加、复制、迭代等

func main() {

	// I. Modify slice elements
	fmt.Println("======= I. Modify slice elements ========")
	a := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	slice1 := a[1:]
	slice2 := a[3:]

	fmt.Println("------- Before Modifications -------")
	fmt.Println("a  = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	slice1[0] = "TUE"
	slice1[1] = "WED"
	slice1[2] = "THU"

	slice2[1] = "FRIDAY"

	fmt.Println("\n-------- After Modifications --------")
	fmt.Println("a  = ", a)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	// II. Length and Capacity
	fmt.Println("\n======= II. Length and Capacity ========")
	/*
		The length of the slice is the number of elements in the slice.
		The capacity is the number of elements in the underlying array starting from the first element in the slice.
	*/
	b := [6]int{10, 20, 30, 40, 50, 60}
	s := b[1:4]

	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s)) // 3, 5

	// III. Extend Capacity
	fmt.Println("\n======= III. Extend Capacity ========")
	s1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original Slice")
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1)) // 10, 10

	// cap() 可以测量切片最长可以达到多少：它等于切片的长度 + 数组除切片之外的长度
	s1 = s1[1:5]
	fmt.Println("\nAfter slicing from index 1 to 5")
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1)) // 4, 9

	// 多个切片如果表示同一个数组的片段，它们可以共享数据；因此一个切片和相关数组的其他切片是共享存储的
	s1 = s1[:8]
	fmt.Println("\nAfter extending the length [:8]")
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1))

	s1 = s1[2:]
	fmt.Println("\nAfter dropping the first two elements")
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1))

	var s2 = s1[2:]
	fmt.Println("\nAfter dropping the first two elements")
	fmt.Printf("s2 = %v, len = %d, cap = %d\n", s2, len(s2), cap(s2))

	// IV. Copy Slices
	fmt.Println("\n======= IV. Copy Slices ========")
	/*
		The copy() function copies elements from one slice to another
		func copy(dst, src []T) int
	*/

	src := []string{"Sublime", "VSCode", "IntelliJ", "Eclipse"}
	dest := make([]string, 2)

	numElementsCopied := copy(dest, src)

	src[0] = "Anna"
	fmt.Println("src = ", src, "capacity = ", cap(src))
	fmt.Println("dest = ", dest, "capacity = ", cap(dest))
	fmt.Println("Number of elements copied from src to dest = ", numElementsCopied)

	// V. Append Elements
	fmt.Println("\n======= V. Append Elements ========")
	// The append() function appends new elements at the end of a given slice.

	// Appending to a slice that doesn't have enough capacity to accommodate new elements
	slice3 := []string{"C", "C++", "Java"}
	slice4 := append(slice3, "Python", "Ruby", "Go")

	fmt.Printf("slice3 = %v, len = %d, cap = %d\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4 = %v, len = %d, cap = %d\n", slice4, len(slice4), cap(slice4))

	slice3[0] = "C#"
	fmt.Println("\nslice3 = ", slice3) // [C# C++ Java]
	fmt.Println("slice4 = ", slice4)   // [C C++ Java Python Ruby Go]

	slicex := slice4[1:3]
	slicexx := append(slicex, "Rust")
	fmt.Println("slicex = ", slicex)                // [C++ Java]
	fmt.Println("slicexx = ", slicexx)              // [C++ Java Rust]
	fmt.Println("slice4 Change to !!!! = ", slice4) // [C C++ Java Rust Ruby Go]

	/*
		In the above example, since slice3 has capacity 3, it can't accommodate more elements.
		So a new underlying array is allocated with bigger capacity when we append more elements to it.
		So if you modify slice3, slice4 won't see those changes because it refers to a different array.
	*/

	// Appending to a slice that has enough capacity to accommodate new elements
	sliceOld := make([]string, 3, 5) // type, length, capacity
	copy(sliceOld, []string{"C", "C++", "Java"})

	// 当使用 append() 向切片添加元素且容量不足时，Go 会自动扩容。
	sliceNew := append(sliceOld, "Python", "Ruby", "Go")

	fmt.Printf("sliceOld = %v, len = %d, cap = %d\n", sliceOld, len(sliceOld), cap(sliceOld))
	fmt.Printf("sliceNew = %v, len = %d, cap = %d\n", sliceNew, len(sliceNew), cap(sliceNew))

	sliceOld[0] = "C#"
	fmt.Println("sliceOld = ", sliceOld)
	fmt.Println("sliceNew = ", sliceNew)

	/*
		In this case, no new array would be allocated, and the elements would be added to the same underlying array.
	*/

	// VI. Append to nil slice
	fmt.Println("\n======= VI. Append to nil slice ========")
	var s3 []string

	// Appending to a nil slice
	s3 = append(s3, "Cat", "Dog", "Lion", "Tiger")

	fmt.Printf("s3 = %v, len = %d, cap = %d\n", s3, len(s3), cap(s3))

	// VII. Append Another Slice
	fmt.Println("\n======= VII. Append Another Slice ========")
	sliceA := []string{"Jack", "John", "Peter"}
	sliceB := []string{"Bill", "Mark", "Steve"}

	sliceC := append(sliceA, sliceB...)

	fmt.Println("sliceA = ", sliceA)
	fmt.Println("sliceB = ", sliceB)
	fmt.Println("After appending sliceA & sliceB = ", sliceC) // [Jack John Peter Bill Mark Steve]

	// VIII. Iterate using for loop
	fmt.Println("\n======= VIII. Iterate using for loop ========")
	countries := []string{"India", "America", "Russia", "England"}

	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}

	// IX. Iterate using range
	fmt.Println("\n======= IX. Iterate using range ========")
	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	for index, number := range primeNumbers {
		fmt.Printf("PrimeNumber(%d) = %d\n", index+1, number)
	}

	// X. Iterate using range with blank identifier
	fmt.Println("\n======= X. Iterate using range with blank identifier ========")
	numbers := []float64{3.5, 7.4, 9.2, 5.4}

	sum := 0.0
	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Total Sum = %.2f\n", sum)
}
