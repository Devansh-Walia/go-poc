package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 10000
	var investmentRate float64 = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow((1+investmentRate/100), years)

	fmt.Println(futureValue)
}
