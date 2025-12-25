package main

import "fmt"

// 切片是一个 长度可变的数组。

func main() {

	// I. Declaring a slice
	fmt.Println("======= I. Declaring a slice ========")

	// var identifier []type
	// Creating a slice using a slice literal
	var s = []int{3, 5, 7, 9, 11, 13, 17} // Creates an array, and returns a slice reference to the array

	// Short hand declaration
	t := []int{2, 4, 8, 16, 32, 64}

	fmt.Println("s = ", s)
	fmt.Println("t = ", t)

	// II. Creating a slice using the make() function
	fmt.Println("======= II. Creating a slice using the make()  ========")
	var slice1 []int = make([]int, 10)
	// load the array/slice:
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}

	// III. print the slice:
	fmt.Println("======= III. print the slice: ========")
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	// IV. Create a slice from array
	fmt.Println("======= IV. Create a slice from array ========")
	var a = [5]string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}

	// Creating a slice from the array
	var aslice []string = a[1:4]

	fmt.Println("Array a = ", a)
	fmt.Println("Slice aslice = ", aslice)

	/*
		low and high parameters are optional in a[low:high]
		The default value for low is 0, and high is the length of the slice.
	*/
	aslice1 := a[1:4]
	aslice2 := a[:3]
	aslice3 := a[2:]
	aslice4 := a[:]

	fmt.Println("aslice1 = ", aslice1)
	fmt.Println("aslice2 = ", aslice2)
	fmt.Println("aslice3 = ", aslice3)
	fmt.Println("aslice4 = ", aslice4)

	// V. Create a slice from another slice
	fmt.Println("======= V. Create a slice from another slice ========")
	cities := []string{"New York", "London", "Chicago", "Beijing", "Delhi", "Mumbai", "Bangalore", "Hyderabad", "Hong Kong"}

	asianCities := cities[3:]
	indianCities := asianCities[1:5]

	fmt.Println("cities = ", cities)
	fmt.Println("asianCities = ", asianCities)
	fmt.Println("indianCities = ", indianCities)

	// VI. The zero value of a slice is nil
	fmt.Println("======= VI. The zero value of a slice is nil ========")

	// A nil slice doesn’t have any underlying array, and has a length and capacity of 0
	var nilSlice []int
	fmt.Printf("nilSlice = %v, len = %d, cap = %d\n", nilSlice, len(nilSlice), cap(nilSlice))

	if nilSlice == nil {
		fmt.Println("nilSlice is nil")
	}
}
