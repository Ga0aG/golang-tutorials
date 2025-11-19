package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main() %d", time.Now().Unix())
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main() %d", time.Now().Unix())
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main() %d", time.Now().Unix())
}

func longWait() {
	fmt.Println("Beginning longWait() %d", time.Now().Unix())
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait() %d", time.Now().Unix())
}

func shortWait() {
	fmt.Println("Beginning shortWait() %d", time.Now().Unix())
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait() %d", time.Now().Unix())
}