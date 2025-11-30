package main

import (
	"fmt"
	"go-demo1/3-struct/bins"
	"go-demo1/3-struct/file"
	"go-demo1/3-struct/storage"
)

func main() {
	fmt.Println("Приложение для работы с контейнерами и мусорными баками")
	storage := storage.NewStorage(file.NewFileStore("bins.json"))
Menu:
	for {
		showMenu()
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}
		switch choice {
		case 1:
			bin := bins.AddNewBin()
			err := storage.AddBin(bin)
			if err != nil {
				fmt.Println("Ошибка:", err)
			}
		case 2:
			storage.ShowBinList()
		case 3:
			break Menu
		default:
			fmt.Println("Некорректный выбор. Пожалуйста, выберите 1 или 2.")
		}
	}
}

func showMenu() {
	fmt.Println("1. Ввести новый Bin")
	fmt.Println("2. Показать все Bin")
	fmt.Println("3. Выход")
	fmt.Print("Выберите действие: ")
}
