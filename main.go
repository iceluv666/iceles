package main

import (
	"fmt"

	"github.com/iceluv666/iceles/greet"
)

func main() {
	var name string
	var table int
	fmt.Printf("Введите ваше имя:")
	fmt.Scan(&name)
	fmt.Printf("Введите номер столика:")
	fmt.Scan(&table)

	greeting := greet.Greet(name, table)

	fmt.Println(greeting)

}
