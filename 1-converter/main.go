package main

import "fmt"

const USDToEURRate = 0.85
const USDToRUBRate = 80.0
const EURToRUBRate = USDToRUBRate / USDToEURRate

func main() {}

func getUserInput() (string, string, float64) {
	var inCurrency string
	var outCurrency string
	var amount float64

	fmt.Print("Введите валюту, из которой хотите конвертировать (USD, EUR, RUB): ")
	fmt.Scan(&inCurrency)
	fmt.Print("Введите валюту, в которую хотите конвертировать (USD, EUR, RUB): ")
	fmt.Scan(&outCurrency)
	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&amount)
	return inCurrency, outCurrency, amount
}

func convertCurrency(inCurrency, outCurrency string, amount float64) float64 {}
