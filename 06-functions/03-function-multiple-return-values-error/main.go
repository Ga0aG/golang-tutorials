package main

import (
	"errors"
	"fmt"
	"math"
)

func getStockPriceChangeWithError(prevPrice, currentPrice float64) (float64, float64, error) {
	if prevPrice == 0 {
		err := errors.New("Previous price cannot be zero")
		return 0, 0, err
	}
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange, nil
}

func main() {
	prevStockPrice := 0.0
	currentStockPrice := 100000.0

	change, percentChange, err := getStockPriceChangeWithError(prevStockPrice, currentStockPrice)
	/*
	Go 语言的函数经常使用两个返回值来表示执行是否成功：返回某个值以及 true 表示成功；返回零值（或 nil）和 false 表示失败（第 4.4 节）。当不使用 true 或 false 的时候，也可以使用一个 error 类型的变量来代替作为第二个返回值：成功执行的话，error 的值为 nil，否则就会包含相应的错误信息（Go 语言中的错误类型为 error: var err error
	*/
	if err != nil {
		fmt.Println("Sorry! There was an error: ", err)
	} else {
		if change < 0 {
			fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
		} else {
			fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
		}
	}
}
