package main

import (
	"fmt"
	"math"
)

const inflationRate = 4.5

func main() {
	var investmentAmount, years, investmentRate float64

	fmt.Print("Enter the amount you want to invest: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the rate you want to invest: ")
	fmt.Scan(&investmentRate)
	fmt.Print("Enter the no of years you want to invest: ")
	fmt.Scan(&years)

	futureValue, inflationAdjustedFutureValue, profit := calculation(investmentAmount, years, investmentRate)

	fmt.Println("future value -> ", futureValue)
	fmt.Println("inflation adjusted future value ->", inflationAdjustedFutureValue)

	fmt.Printf("Total profit of %.2f was achieved", profit)
}

func calculation(investmentAmount, years, investmentRate float64) (float64, float64, float64) {
	fv := investmentAmount * math.Pow((1+investmentRate/100), years)
	inv := fv / math.Pow(inflationRate/100, years)
	profit := inv - investmentAmount

	return fv, inv, profit
}
