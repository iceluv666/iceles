package greet

import (
	"fmt"
)

func Greet(name string, table int) string {
	return fmt.Sprintf("Привет, %s! Your table is %d.", name, table)
}
