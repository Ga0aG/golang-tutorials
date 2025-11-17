package main

import "fmt"
import "strings"

func main() {
	// Normal String (Can not contain newlines, and can have escape characters like `\n`, `\t` etc)
	var website = "\thttps://www.callicoder.com\t\n"

	// Raw String (Can span multiple lines. Escape characters are not interpreted)
	var siteDescription = `\t\tCalliCoder is a programming blog where you can find
                           practical guides and tutorials on programming languages, 
                           web development, and desktop app development.\t\n`

	fmt.Println(website, siteDescription)

	// 前缀和后缀
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? have suffix %s", str, "Th", "string")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("%t\n", strings.HasSuffix(str, "string"))

  // 索引
	str = "Hi, I'm Marc, Hi."
	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))

	// 其他方法
	// Contains, Count, Replace, Repeat, ToLower, ToUpper, Split
}
