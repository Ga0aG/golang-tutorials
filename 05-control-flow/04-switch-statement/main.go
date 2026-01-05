package main
import "fmt"

func main() {
	// ===== Switch Statement =====
	var dayOfWeek = 6
	switch dayOfWeek {
		case 1: fmt.Println("Monday")
		case 2: fmt.Println("Tuesday")
		case 3: fmt.Println("Wednesday")
		case 4: fmt.Println("Thursday")
		case 5: fmt.Println("Friday")
		case 6: {
			fmt.Println("Saturday")
			fmt.Println("Weekend. Yaay!")
		}
		case 7: {
			fmt.Println("Sunday")
			fmt.Println("Weekend. Yaay!")
		}
		default: fmt.Println("Invalid day")
	}

	// ===== Switch with Combined Cases =====
	switch dayOfWeek := 5; dayOfWeek {
		case 1, 2, 3, 4, 5:
			fmt.Println("Weekday")
		case 6, 7:
			fmt.Println("Weekend")
		default:
			fmt.Println("Invalid Day")
	}

	// ===== Switch with No Expression (like if-else) =====
	var BMI = 21.0
	switch {
		case BMI < 18.5:
			fmt.Println("You're underweight")
		case BMI >= 18.5 && BMI < 25.0:
			fmt.Println("Your weight is normal")
		case BMI >= 25.0 && BMI < 30.0:
			fmt.Println("You're overweight")
		default:
			fmt.Println("You're obese")
	}
}