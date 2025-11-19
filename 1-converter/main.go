package main

import (
	"fmt"
	"strings"
)

var currencyRates = func() *map[string]map[string]float64 {
	m := map[string]map[string]float64{
		"USD": map[string]float64{
			"EUR": 0.85,
			"RUB": 80.0,
		},
		"EUR": map[string]float64{
			"USD": 1.18,
			"RUB": 94.0,
		},
		"RUB": map[string]float64{
			"USD": 0.0125,
			"EUR": 0.0106,
		},
	}
	return &m
}()

func main() {
	fmt.Println("Конвертер валют")
	for {
		inCurrency, outCurrency, amount := getUserInput(currencyRates)
		convertedAmount := convertCurrency(currencyRates, inCurrency, outCurrency, amount)
		fmt.Printf("Конвертированная сумма: %.2f %s\n", convertedAmount, outCurrency)
		isFinish := askFinish()
		if !isFinish {
			break
		}
	}
}

func getUserInput(rates *map[string]map[string]float64) (string, string, float64) {
	var inCurrency string
	var outCurrency string
	var amount float64

	fmt.Print("Введите валюту, из которой хотите конвертировать (USD, EUR, RUB): ")
	fmt.Scan(&inCurrency)

	if _, ok := (*rates)[inCurrency]; !ok {
		fmt.Println("Неверная валюта. Попробуйте снова.")
		return getUserInput(rates)
	}

	outCurrency = getOutCurrency(rates, inCurrency)

	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Сумма должна быть положительным числом. Попробуйте снова.")
		return getUserInput(rates)
	}
	return inCurrency, outCurrency, amount
}

func convertCurrency(rates *map[string]map[string]float64, inCurrency, outCurrency string, amount float64) float64 {
	r, ok := (*rates)[inCurrency]
	if !ok {
		panic("внутренняя ошибка: неизвестная исходная валюта")
	}
	rate, ok := r[outCurrency]
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

func getOutCurrency(rates *map[string]map[string]float64, inCurrency string) string {
	var outCurrency string
	inner := (*rates)[inCurrency]
	targets := keys(&inner)
	fmt.Printf("Введите валюту, в которую хотите конвертировать (%s): ", strings.Join(targets, ", "))
	fmt.Scan(&outCurrency)

	if _, ok := (*rates)[inCurrency][outCurrency]; !ok {
		fmt.Println("Неверная валюта. Попробуйте снова.")
		return getOutCurrency(rates, inCurrency)
	}
	return outCurrency
}

func keys(m *map[string]float64) []string {
	res := make([]string, 0, len(*m))
	for k := range *m {
		res = append(res, k)
	}
	return res
}
