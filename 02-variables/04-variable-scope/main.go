package main

var a = "G"
var b = "M"

func main() {
	print("Local scope\n")
	n()
	m()
	n()
	print("\nGlobal scope\n")
	n1()
	m1()
	n1()
	// GOGMOO
}

func n() { print(a) }

func m() {
	a := "O"
	print(a)
}

func n1() { print(b) }

func m1() {
	b = "O"
	print(b)
}
