package greet

import (
	"testing"
)

func TestingGreet(t *testing.T) {
	name := "Alesha"
	talble := 5
	expected := "Добро пожаловать, Алеша! Ваш столик под номером 5"
	resilt := Greet(name, talble)

	if resilt != expected {
		t.Errorf("Ожидал приветствие %s, но получил %s", expected, resilt)
	}
}
