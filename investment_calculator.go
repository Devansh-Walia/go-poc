package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 4.5
	investmentAmount, years, investmentRate := 10000.0, 10.0, 5.5

	futureValue := investmentAmount * math.Pow((1+investmentRate/100), years)
	inflationAdjustedFutureValue := futureValue / math.Pow(inflationRate/100, years)

	fmt.Println("future value -> ", futureValue)
	fmt.Println("inflation adjusted future value ->", inflationAdjustedFutureValue)
}
