package main

import (
	"fmt"

	"github.com/iceluv666/iceles/greet"
)

func main() {
	text := greet.Greet("Sasha", 5)

	fmt.Println(text)
}
