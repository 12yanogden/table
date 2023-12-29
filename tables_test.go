package tables

import (
	"reflect"
	"testing"
)

func TestSlices(t *testing.T) {
	table := "1:1\t1:2\t1:3\n2:1\t2:2\t2:3\n3:1\t3:2\t3:3"
	expected := [][]string{
		{"1:1", "1:2", "1:3"},
		{"2:1", "2:2", "2:3"},
		{"3:1", "3:2", "3:3"},
	}
	actual := Slices(table)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\n" + Table(expected) + "\nActual:\n" + Table(actual))
	}
}
