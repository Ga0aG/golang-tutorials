package main

import "fmt"

func main() {
    fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
    // var i1 int = MultiPly3Nums(2, 5, 6)
    // fmt.Printf("MultiPly 2 * 5 * 6 = %d\n", i1)
    a := 2
    b := &a
    CallByPointer(b)
    fmt.Printf("a after CallByPointer: %d\n", a)
}

// Go 默认使用按值传递来传递参数
func MultiPly3Nums(a int, b int, c int) int {
    // var product int = a * b * c
    // return product
    return a * b * c
}

func CallByPointer(ptr *int) {
	*ptr = 1
}