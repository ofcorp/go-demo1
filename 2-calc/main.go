package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main() {

	operationsMap := map[string]func([]int){
	"AVG": func(numbers []int) {
		sum := sumNumbers(numbers)
		avg := float64(sum) / float64(len(numbers))
		fmt.Printf("Среднее значение: %.2f\n", avg)
	},
	"SUM": func(numbers []int) {
		sum := sumNumbers(numbers)
		fmt.Printf("Сумма: %d\n", sum)
	},
	"MED": func(numbers []int) {
		sorted := make([]int, len(numbers))
		copy(sorted, numbers)
		for i := 0; i < len(sorted)-1; i++ {
			for j := 0; j < len(sorted)-i-1; j++ {
				if sorted[j] > sorted[j+1] {
					sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
				}
			}
		}
		var median float64
		mid := len(sorted) / 2
		if len(sorted)%2 == 0 {
			median = float64(sorted[mid-1]+sorted[mid]) / 2.0
		} else {
			median = float64(sorted[mid])
		}
		fmt.Printf("Медиана: %.2f\n", median)
	},
}

	fmt.Println("Калькулятор")
	operation := userInputOperation()
	numbers := userInputNumbers()
	operationsMap[operation](numbers)
}

func userInputOperation() string {
	var operation string
	println("Введите операцию (AVG, SUM, MED): ")
	fmt.Scan(&operation)
	switch operation {
	case "AVG", "SUM", "MED":
		return operation
	default:
		fmt.Println("Неверная операция. Попробуйте снова.")
		return userInputOperation()
	}
}

func userInputNumbers() []int {
	fmt.Print("Введите числа через запятую: ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, ",")
	var nums []int
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			fmt.Println("Ошибка парсинга числа:", p)
			return userInputNumbers()
		}
		nums = append(nums, n)
	}
	if len(nums) == 0 {
		fmt.Println("Не введено ни одного числа. Повторите.")
		return userInputNumbers()
	}
	return nums
}

func sumNumbers(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}
