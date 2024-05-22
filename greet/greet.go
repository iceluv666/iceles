package greet

import (
   "fmt"
)

func Greet(name string, table int) string{
	return fmt.Sprintf("Добро пожаловать! %s Ваш столик %d", name, table)

}