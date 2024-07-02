package main

import (
	"fmt"
	"math"
)

func compoundInterest(principal float64, rate float64, time float64, n float64) float64 {
	amount := principal * math.Pow((1+rate/(n*100)), n*time)
	interest := amount - principal
	return interest
}

func main() {
	var principal, rate, time, n float64

	fmt.Print("Enter the value of principal ")
	fmt.Scan(&principal)
	fmt.Print("Enter the value of rate(in %)")
	fmt.Scan(&rate)
	fmt.Print("enter the value of time")
	fmt.Scan(&time)
	fmt.Print("enter the value of n ")
	fmt.Scan(&n)

	interest := compoundInterest(principal, rate, time, n)

	fmt.Print(interest)
}
