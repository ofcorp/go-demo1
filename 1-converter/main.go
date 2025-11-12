package main

import "fmt"

const USDToEURRate = 0.85
const USDToRUBRate = 80.0
const EURToRUBRate = USDToRUBRate / USDToEURRate

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
	switch inCurrency {
	case "USD", "EUR", "RUB":
	default:
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
	switch inCurrency {
	case "USD":
		if outCurrency == "EUR" {
			return amount * USDToEURRate
		} else if outCurrency == "RUB" {
			return amount * USDToRUBRate
		}
	case "EUR":
		if outCurrency == "USD" {
			return amount / USDToEURRate
		} else if outCurrency == "RUB" {
			return amount * EURToRUBRate
		}
	case "RUB":
		if outCurrency == "USD" {
			return amount / USDToRUBRate
		} else if outCurrency == "EUR" {
			return amount / EURToRUBRate
		}
	}
	panic("внутренняя ошибка")
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
	fmt.Print("Введите валюту, в которую хотите конвертировать ")
	switch inCurrency {
	case "USD":
		fmt.Print("(EUR, RUB): ")
		fmt.Scan(&outCurrency)
		if outCurrency != "EUR" && outCurrency != "RUB" {
			fmt.Println("Неверная валюта. Попробуйте снова.")
			return getOutCurrency(inCurrency)
		}
	case "EUR":
		fmt.Print("(RUB, USD): ")
		fmt.Scan(&outCurrency)
		if outCurrency != "RUB" && outCurrency != "USD" {
			fmt.Println("Неверная валюта. Попробуйте снова.")
			return getOutCurrency(inCurrency)
		}
	case "RUB":
		fmt.Print("(USD, EUR): ")
		fmt.Scan(&outCurrency)
		if outCurrency != "USD" && outCurrency != "EUR" {
			fmt.Println("Неверная валюта. Попробуйте снова.")
			return getOutCurrency(inCurrency)
		}
	default:
		panic("внутренняя ошибка")
	}
	return outCurrency
}
