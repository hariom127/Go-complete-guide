package main

import (
	"fmt"
	"math"
)

func main() {
	var inflationRate float64
	var amount float64 //amount is by default set to null value (its a float64 that why the default null value is 0.0)
	var returnRate float64
	var years float64

	fmt.Print("Enter investment amount: ")
	//take value as input from user using Scan
	fmt.Scan(&amount)

	fmt.Print("Enter returnRate: ")
	fmt.Scan(&returnRate)

	fmt.Print("Enter inflation rate: ")
	fmt.Scan(&inflationRate)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureAndRealValue(amount, returnRate, inflationRate, years)

	// var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Printf("\n Future Value: %.0f\n", futureValue) // after decimal zero value
	// fmt.Printf("\n Future Value: %.2f\n", futureValue) // after decimal two value
	fmt.Printf("\n Future Value: %.2f\n futureRealValue Value: %.3f \n\n", futureValue, futureRealValue) // multi value in print f
	// fmt.Printf("\n Future Value: %v\n", futureValue)
	// fmt.Println("Future Value: ", futureValue)
	fmt.Println("futureRealValue Value: ", futureRealValue)

}

func calculateFutureAndRealValue(amount, returnRate, years, inflationRate float64) (fv float64, rv float64) {
	fv = amount * math.Pow(1+returnRate/100, years)
	rv = fv / math.Pow(1+inflationRate/100, years)
	// return
	return fv, rv
}
