package tables

import (
	"strconv"
	"strings"

	"github.com/12yanogden/shell"
	"github.com/12yanogden/strslices"
)

func Slices(table string) [][]string {
	rows := strslices.Explode(table, "\n")
	var slices [][]string

	for _, row := range rows {
		slices = append(slices, strings.Fields(row))
	}

	return slices
}

func Table(slices [][]string) string {
	TAB_WIDTH := 8
	columnWidths := getColumnWidths(slices, TAB_WIDTH)
	table := ""

	// Set tab width in terminal manually
	shell.Run("tabs", []string{"-" + strconv.Itoa(TAB_WIDTH)})

	// Build table
	for _, row := range slices {
		for colIndex, value := range row {
			tabCount := calcTabCount(value, columnWidths[colIndex], TAB_WIDTH)

			table += (value + strings.Repeat("\t", tabCount))
		}

		table += "\n"
	}

	return table
}

func getColumnWidths(slices [][]string, TAB_WIDTH int) []int {
	columnWidths := []int{}

	// Get the maximum length of the values by column
	for _, slice := range slices {
		for i, value := range slice {
			if i > (len(columnWidths) - 1) {
				columnWidths = append(columnWidths, len(value))
			} else if len(value) > columnWidths[i] {
				columnWidths[i] = len(value)
			}
		}
	}

	// Pad columns for space between
	for i, width := range columnWidths {
		remainder := width % TAB_WIDTH

		if remainder == 0 {
			remainder = TAB_WIDTH
		}

		columnWidths[i] = width + remainder
	}

	return columnWidths
}

func calcTabCount(value string, colWidth int, TAB_WIDTH int) int {
	padding := (colWidth - len(value))
	tabCount := padding / TAB_WIDTH

	if (padding % TAB_WIDTH) > 0 {
		tabCount++
	}

	return tabCount
}
