package main

import (
	"fmt"
)

func main() {
	callback(1, Add)
	callback2(1, Add2)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func Add2(a, b int) int {
	c := a+b
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, c)
	return c
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

func callback2(y int, f func(int, int) int) {
	f(y, 2) // this becomes Add(1, 2)
}