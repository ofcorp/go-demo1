package main

import "fmt"

const USDToEURRate = 0.85
const USDToRURRate = 80.0

func main() {
	amountInEUR := 100.0
	EURtoRUR := amountInEUR / USDToEURRate * USDToRURRate
	fmt.Printf("%.2f EUR равно %.2f RUR\n", amountInEUR, EURtoRUR)
}
