package main

import "fmt"

func main() {

LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	fmt.Printf("======================\n")

	for i2 := 0; i2 <= 5; i2++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				goto LABEL2
			}
			fmt.Printf("i is: %d, and j is: %d\n", i2, j)
		}
	}
LABEL2:
	fmt.Printf("Jump here")

}
