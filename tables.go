package tables

import "strings"

func ToSlices(table string) [][]string {
	rows := strings.Split(table, "\n")
	var slices [][]string

	for _, row := range rows {
		slices = append(slices, strings.Fields(row))
	}

	return slices
}
