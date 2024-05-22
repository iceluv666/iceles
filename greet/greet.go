package greet

import (
	"fmt"
)

func Greet(name string, table int) string {
	return fmt.Sprintf("Добро пожаловать, %s! Ваш столик под номером %d", name, table)

}
