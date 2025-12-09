package main

import "fmt"

func main() {
    a := 2
    b := &a
    CallByPointer(b)
    fmt.Printf("a after CallByPointer: %d\n", a)
}

func CallByPointer(ptr *int) {
	*ptr = 1
}