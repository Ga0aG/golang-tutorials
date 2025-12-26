package main

import (
	"fmt"
	"math"
	"strings"
)

// Go 基础数据类型：数值类型、字符、布尔值、复数、字符串

func main() {

	// I. Numeric Types
	fmt.Println("======= I. Numeric Types ========")

	var myInt8 int8 = 97

	/*
	  When you don't declare any type explicitly, the type inferred is `int`
	  (The default type for integers)
	*/
	var myInt = 1200

	var myUint uint = 500

	var myHexNumber = 0xFF  // Use prefix '0x' or '0X' for declaring hexadecimal numbers
	var myOctalNumber = 034 // Use prefix '0' for declaring octal numbers

	var myFloat32 float32 = 4.5
	var myFloat = 9.12 // Type inferred as `float64` (the default type for floating-point numbers)

	fmt.Printf("%d, %d, %d, %#x, %#o %f %f\n", myInt8, myInt, myUint, myHexNumber, myOctalNumber, myFloat32, myFloat)

	// II. Characters
	fmt.Println("\n======= II. Characters ========")

	var myByte byte = 'a' // byte 类型是 uint8 的别名
	var myRune rune = '♥' // unicode, 是 int32 的别名

	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)

	// III. Numeric Operations
	fmt.Println("\n======= III. Numeric Operations ========")

	var a, b = 4, 5
	var res1 = (a + b) * (a + b) / 2 // Arithmetic operations

	a++ // Increment a by 1

	b += 10 // Increment b by 10

	var res2 = a ^ b // Bitwise XOR

	var r = 3.5
	var res3 = math.Pi * r * r // Operations on floating-point type

	fmt.Printf("res1 : %v, res2 : %v, res3 : %v\n", res1, res2, res3)

	// IV. Boolean Types
	fmt.Println("\n======= IV. Boolean Types ========")

	var myBoolean bool = true
	var anotherBoolean = false // Inferred type

	var truth = 3 <= 5
	var falsehood = 10 != 10

	// Short Circuiting
	var res4 = 10 > 20 && 5 == 5     // Second operand is not evaluated since first evaluates to false
	var res5 = 2*2 == 4 || 10%3 == 0 // Second operand is not evaluated since first evaluates to true

	fmt.Println(myBoolean, anotherBoolean, truth, falsehood, res4, res5)

	// V. Complex Numbers
	fmt.Println("\n======= V. Complex Numbers ========")

	// === Creating complex numbers ====
	/*
		complex64: both real and imaginary parts are of float32 type.
		complex128: both real and imaginary parts are of float64 type.
	*/
	var x complex64 = 3.4 + 2.9i
	var y = 5 + 7i // Type inferred as `complex128` (default type for complex numbers)

	fmt.Println(x, y)

	// Creating complex no from variables
	var c1 = 4.5
	var c2 = 7.1

	var c = complex(c1, c2) // c1 + c2*i won't work
	fmt.Println(c)

	// ===== Complex No Operations =====
	var comp1 = 3 + 5i
	var comp2 = 2 + 4i

	var compRes1 = comp1 + comp2
	var compRes2 = comp1 - comp2
	var compRes3 = comp1 * comp2
	var compRes4 = comp1 / comp2

	fmt.Println(compRes1, compRes2, compRes3, compRes4) // (5+9i) (1+1i) (-14+22i) (1.3-0.1i)

	// VI. Strings
	fmt.Println("\n======= VI. Strings ========")

	// Normal String (Can not contain newlines, and can have escape characters like `\n`, `\t` etc)
	var website = "\thttps://www.callicoder.com\t\n"

	// Raw String (Can span multiple lines. Escape characters are not interpreted)
	var siteDescription = `\t\tCalliCoder is a programming blog where you can find
                           practical guides and tutorials on programming languages,
                           web development, and desktop app development.\t\n`

	fmt.Println(website, siteDescription)

	// 前缀和后缀
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? have suffix %s\n", str, "Th", "string")
	fmt.Printf("HasPrefix: %t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("HasSuffix: %t\n", strings.HasSuffix(str, "string"))

	// 索引
	str = "Hi, I'm Marc, Hi."
	fmt.Printf("\nThe position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))

	// 其他方法: Contains, Count, Replace, Repeat, ToLower, ToUpper, Split
}
