package main

import "fmt"

func main() {
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
}
