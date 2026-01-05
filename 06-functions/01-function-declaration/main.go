package main

import (
	"errors"
	"fmt"
	"math"
)

// ===== Basic Function Declaration =====
func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

// ===== Function with Multiple Return Values =====
// 当多个连续参数类型相同时，可以只在最后一个参数声明类型
func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange
}

// ===== Function with Multiple Return Values + Error =====
func getStockPriceChangeWithError(prevPrice, currentPrice float64) (float64, float64, error) {
	if prevPrice == 0 {
		err := errors.New("Previous price cannot be zero")
		return 0, 0, err
	}
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange, nil
}

// ===== Function with Named Return Values =====
func getNamedStockPriceChange(prevPrice, currentPrice float64) (change, percentChange float64) { // 在函数签名中，你可以给返回值起名字，这样这些变量在函数体内就可以直接使用，无需再声明
	change = currentPrice - prevPrice // change就不用change := ...了
	percentChange = (change / prevPrice) * 100
	return change, percentChange
}

// ===== Call By Value (默认) =====
func MultiPly3Nums(a int, b int, c int) int {
	return a * b * c
}

// ===== Call By Pointer =====
func CallByPointer(ptr *int) {
	*ptr = 1
}

func main() {
	// ===== Basic Function Call =====
	fmt.Println("===== Basic Function Call =====")
	x := 5.75
	y := 6.25
	result := avg(x, y)
	fmt.Printf("Average of %.2f and %.2f = %.2f\n", x, y, result)

	// ===== Multiple Return Values =====
	fmt.Println("\n===== Multiple Return Values =====")
	prevStockPrice := 0.0
	currentStockPrice := 100000.0
	change, percentChange := getStockPriceChange(prevStockPrice, currentStockPrice)
	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
	}

	// ===== Multiple Return Values with Error =====
	fmt.Println("\n===== Multiple Return Values with Error =====")
	prevStockPrice = 0.0
	currentStockPrice = 100000.0
	change, percentChange, err := getStockPriceChangeWithError(prevStockPrice, currentStockPrice)
	if err != nil {
		fmt.Println("Sorry! There was an error: ", err)
	} else {
		if change < 0 {
			fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
		} else {
			fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
		}
	}

	// ===== Named Return Values =====
	fmt.Println("\n===== Named Return Values =====")
	prevStockPrice = 100000.0
	currentStockPrice = 90000.0
	change, percentChange = getNamedStockPriceChange(prevStockPrice, currentStockPrice)
	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
	}

	// ===== Blank Identifier (ignoring multiple return values) =====
	fmt.Println("\n===== Blank Identifier =====")
	prevStockPrice = 80000.0
	currentStockPrice = 120000.0
	change, _ = getStockPriceChange(prevStockPrice, currentStockPrice) // 使用空白标识符 _ 来忽略不需要的返回值
	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f\n", math.Abs(change))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f\n", change)
	}

	// ===== Call By Value vs Call By Pointer =====
	fmt.Println("\n===== Call By Value vs Call By Pointer =====")
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
	a := 2
	fmt.Printf("a before CallByPointer: %d\n", a)
	b := &a
	CallByPointer(b)
	fmt.Printf("a after CallByPointer: %d\n", a)
}
