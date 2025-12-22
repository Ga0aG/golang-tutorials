package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 值接收者方法
func (p Person) GetName() string {
	return p.Name
}

// 指针接收者方法
func (p *Person) IncreAge() {
	p.Age++
}

type Item struct {
		Value int
}

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d *Dog) Speak() string {
	return "Woof"
}

func main() {
	// I. 访问结构体字段
	p := &Person{Name: "Alice", Age: 30}
	// 以下两种写法等价，Go 会自动解引用：
	fmt.Println((*p).Name) // 显式解引用
	fmt.Println(p.Name)    // 自动解引用 ✅（最常用）

	// 甚至方法调用时：
	p.Name = "Bob"  // 自动解引用并赋值

	// II. 调用方法
	fmt.Println(p.GetName())
	p.IncreAge()
	fmt.Println(p.Age)

	// III. 反面例子，切片、映射、通道的元素访问
	slicePtr := &[]int{1, 2, 3}
	(*slicePtr)[0] = 10  // 显式解引用
	// slicePtr[0] = 10     // 错误！❌ 切片指针不能这样用

	// 但如果是结构体切片：
	items := &[]Item{{1}, {2}, {3}}
	// items[0].Value = 10     // 错误！❌ 不能直接索引
	(*items)[0].Value = 10    // 需要显式解引用

	// IV. 接口
	var dog Dog
	var speaker Speaker
	// Go 会自动将 dog 转换为 &dog 以满足接口
	speaker = dog      // 自动取地址 ✅
	speaker = &dog     // 也可以显式
	speaker.Speak()    // 正常调用

	
}
