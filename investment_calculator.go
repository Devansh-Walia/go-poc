package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 4.5
	var investmentAmount, years, investmentRate float64

	fmt.Print("Enter the amount you want to invest: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the rate you want to invest: ")
	fmt.Scan(&investmentRate)
	fmt.Print("Enter the no of years you want to invest: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+investmentRate/100), years)
	inflationAdjustedFutureValue := futureValue / math.Pow(inflationRate/100, years)
	profit := inflationAdjustedFutureValue - investmentAmount

	fmt.Println("future value -> ", futureValue)
	fmt.Println("inflation adjusted future value ->", inflationAdjustedFutureValue)

	fmt.Printf("Total profit of %.2f was achieved", profit)
}
