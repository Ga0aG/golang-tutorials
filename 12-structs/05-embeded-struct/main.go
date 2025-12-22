package main

import "fmt"
import "math"

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

/*
当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？

外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。没有办法来解决这种问题引起的二义性，必须由程序员自己修正。
*/


type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	name string
}

// 聚合/组合
type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

func main() {
	b := B{A{1, 2}, 3.0, 4.0}
	// 内层结构体被简单的插入或者内嵌进外层结构体。
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
	fmt.Println(b)

	// 调用内嵌结构体的方法
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs()) // 打印 5

	// 组合
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	// shorter
	c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	// fmt.Println(c) &{Barak Obama 1 - Yes we can!}
	c.Log().Add("2 - After me the world will be a better place!")
	//fmt.Println(c.log)
	fmt.Println(c.Log())
}