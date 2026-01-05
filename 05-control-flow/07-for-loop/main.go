package main
import "fmt"

func main() {
	// ===== Basic For Loop =====
	fmt.Println("Basic for loop:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// ===== For Loop Without Init Statement =====
	fmt.Println("\nFor loop without init statement:")
	i := 2
	for ;i <= 10; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// ===== For Loop Without Increment Statement =====
	fmt.Println("\nFor loop without increment statement:")
	i = 2
	for ;i <= 20; {
		fmt.Printf("%d ", i)
		i *= 2
	}
	fmt.Println()

	// ===== Infinite Loop (with break) =====
	fmt.Println("\nInfinite loop with break:")
	count := 0
	for {
		if count >= 3 {
			break
		}
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Println()

	// ===== For Loop with Break =====
	fmt.Println("\nFor loop with break:")
	for num := 1; num <= 100; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Printf("First positive number divisible by both 3 and 5 is %d\n", num)
			break
		}
	}

	// ===== For Loop with Continue =====
	fmt.Println("For loop with continue (print odd numbers):")
	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}