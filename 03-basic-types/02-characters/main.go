package main

import "fmt"

func main() {
	var myByte byte = 'a' // byte 类型是 uint8 的别名
	var myRune rune = '♥' // unicode, 是 int32 的别名

	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)
}