package ascii_table

import (
	"fmt"
	"testing"
)

func TestAddItem(t *testing.T) {
	tb := New()
	tb.AddItem("Name", "Alice").
		AddItem("Name", "Bob").
		AddItem("Name", "Charlie")
	tb.AddItem("Age", "30").
		AddItem("Age", "25").
		AddItem("Age", "35")
	tb.AddItem("City", "New York").
		AddItem("City", "Los Angeles").
		AddItem("City", "Chicago")
	fmt.Println(tb.String())
}

func TestAddRow(t *testing.T) {
	tb := New()
	tb.AddRow(map[string]string{
		"Name": "Alice",
		"Age":  "30",
		"City": "New York",
	}).AddRow(map[string]string{
		"Name": "Bob",
		"Age":  "35",
		"City": "Los Angeles",
	}).AddRow(map[string]string{
		"Name": "Charlie",
		"Age":  "25",
		"City": "New York",
	})
	fmt.Println(tb.String())
}
