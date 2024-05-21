package tables

// Table represents a table in the cafe
type Table struct {
	Location string // "внутри" or "снаружи"
	Number  int    // Table number
	Occupied bool  // Is the table occupied
}

// cafesTables holds all tables for the cafe
var cafesTables = map[string][]Table{
	"внутри": {
		{"внутри", 1, false},
		{"внутри", 2, false},
		{"внутри", 3, false},
		{"внутри", 4, false},
		{"внутри", 5, false},
	},
	"снаружи": {
		{"снаружи", 1, false},
		{"снаружи", 2, false},
		{"снаружи", 3, false},
	},
}

// GetTables returns all tables for a specific location
func GetTables(location string) []Table {
	return cafesTables[location]
}

// ReserveTable attempts to reserve a table for a customer
func ReserveTable(name string, location string, tableNumber int) (bool, string) {
	tables := GetTables(location)
	if tableNumber < 1 || tableNumber > len(tables) {
		return false, fmt.Sprintf("Столика с номером %d не существует.", tableNumber)
	}

	table := tables[tableNumber-1]
	if table.Occupied {
		return false, fmt.Sprintf("Столик №%d (%s) занят. Пожалуйста, выберите другой номер.", tableNumber, location)
	}

	table.Occupied = true
	return true, fmt.Sprintf("Привет, %s! Ваш столик №%d (%s).", name, tableNumber, location)
}
