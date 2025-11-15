package main

import (
	"fmt"
	"strings"
)

var currencyRates = map[string]map[string]float64{
	"USD": {
		"EUR": 0.85,
		"RUB": 80.0,
	},
	"EUR": {
		"USD": 1.18,
		"RUB": 94.0,
	},
	"RUB": {
		"USD": 0.0125,
		"EUR": 0.0106,
	},
}

func main() {
	fmt.Println("Конвертер валют")
	for {
		inCurrency, outCurrency, amount := getUserInput()
		convertedAmount := convertCurrency(inCurrency, outCurrency, amount)
		fmt.Printf("Конвертированная сумма: %.2f %s\n", convertedAmount, outCurrency)
		isFinish := askFinish()
		if !isFinish {
			break
		}
	}
}

func getUserInput() (string, string, float64) {
	var inCurrency string
	var outCurrency string
	var amount float64

	fmt.Print("Введите валюту, из которой хотите конвертировать (USD, EUR, RUB): ")
	fmt.Scan(&inCurrency)

	if _, ok := currencyRates[inCurrency]; !ok {
		fmt.Println("Неверная валюта. Попробуйте снова.")
		return getUserInput()
	}

	outCurrency = getOutCurrency(inCurrency)

	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Сумма должна быть положительным числом. Попробуйте снова.")
		return getUserInput()
	}
	return inCurrency, outCurrency, amount
}

func convertCurrency(inCurrency, outCurrency string, amount float64) float64 {
	rates, ok := currencyRates[inCurrency]
	if !ok {
		panic("внутренняя ошибка: неизвестная исходная валюта")
	}
	rate, ok := rates[outCurrency]
	if !ok {
		panic("внутренняя ошибка: недоступное направление конвертации")
	}
	return amount * rate
}

func askFinish() bool {
	var isFinish string
	fmt.Print("Хотите ещё конвертировать? (y/n): ")
	fmt.Scan(&isFinish)
	if isFinish == "y" || isFinish == "Y" {
		return true
	}
	return false
}

func getOutCurrency(inCurrency string) string {
	var outCurrency string
	targets := keys(currencyRates[inCurrency])
	fmt.Printf("Введите валюту, в которую хотите конвертировать (%s): ", strings.Join(targets, ", "))
	fmt.Scan(&outCurrency)

	if _, ok := currencyRates[inCurrency][outCurrency]; !ok {
		fmt.Println("Неверная валюта. Попробуйте снова.")
		return getOutCurrency(inCurrency)
	}
	return outCurrency
}

func keys(m map[string]float64) []string {
	res := make([]string, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}
