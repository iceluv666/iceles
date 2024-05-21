package main

import (
	"fmt"
	"strconv"

	tables "github.com/iceluv666/iceles/tables"
)

func main() {
	// Get user information
	name, location, tableNumber := getUserInput()

	// Reserve table
	reserved, message := tables.ReserveTable(name, location, tableNumber)

	// Print confirmation message
	fmt.Println(message)

	if reserved {
		fmt.Printf("Столик №%d (%s) забронирован на имя %s.\n", tableNumber, location, name)
	}
}

func getUserInput() (string, string, int) {
	var name, location string
	var tableNumber int

	fmt.Print("Введите ваше имя: ")
	fmt.Scanln(&name)

	for {
		fmt.Print("Выберите место (внутри/снаружи): ")
		fmt.Scanln(&location)

		if location == "внутри" || location == "снаружи" {
			break
		}

		fmt.Println("Неверный ввод. Пожалуйста, введите 'внутри' или 'снаружи'.")
	}

	for {
		fmt.Printf("Выберите номер столика (%s): ", location)
		fmt.Scanln(&input)

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Неверный ввод. Пожалуйста, введите номер столика.")
			continue
		}

		tables := tables.GetTables(location)
		if number < 1 || number > len(tables) {
			fmt.Printf("Столика с номером %d не существует. Пожалуйста, выберите другой номер.\n", number)
			continue
		}

		tableNumber = number
		break
	}

	return name, location, tableNumber
}
