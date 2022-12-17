package prntme

import (
	"strings"

	"github.com/olekukonko/tablewriter"
)

// Takes a title []string and a 2d array of values. Returns a table.
func ReturnasTable(titles []string, values [][]string) string {
	var strbuf strings.Builder

	table := tablewriter.NewWriter(&strbuf)
	table.SetHeader(titles)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(values) // Add Bulk Data
	table.Render()

	tablestr := strbuf.String()

	return tablestr

}
