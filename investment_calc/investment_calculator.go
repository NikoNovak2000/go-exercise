package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	OutputText("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	OutputText("Enter the number of years: ")
	fmt.Scan(&years)

	OutputText("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Formatting
	// formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future Value adjusted for inflation: %.1f\n", futureRealValue)
	// fmt.Print(formattedFV, formattedRFV)

	// fmt.Printf("Future Value: %.1f\nFuture Value adjusted for inflation: %.1f\n", futureValue, futureRealValue)

	fmt.Printf(`Future Value: %.1f
	Future Value adjusted for inflation: %.1f`,
		futureValue, futureRealValue)
}

// Wrapper func
func OutputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
