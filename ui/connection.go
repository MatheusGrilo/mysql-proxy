package ui

import (
	"github.com/rivo/tview"
)

// CreateConnectionContent creates the content for the "Connection" section
func CreateConnectionContent() *tview.Flex {
	table := tview.NewTable().
		SetBorders(false)

	// MySQL - Currently connections
	table.SetCell(0, 0, tview.NewTableCell("[default]MySQL: [yellow]192.168.100.254:3306").
		SetTextColor(tview.Styles.PrimaryTextColor).
		SetAlign(tview.AlignLeft).
		SetExpansion(1))
	table.SetCell(0, 1, tview.NewTableCell("[default]Connections made: [red]1003").
		SetTextColor(tview.Styles.SecondaryTextColor).
		SetAlign(tview.AlignRight).
		SetMaxWidth(30))

	// Proxy - Total connections made
	table.SetCell(1, 0, tview.NewTableCell("[default]Proxy: [green]127.60.0.1:4040").
		SetTextColor(tview.Styles.PrimaryTextColor).
		SetAlign(tview.AlignLeft).
		SetExpansion(1))
	table.SetCell(1, 1, tview.NewTableCell("[default]Currently connections: [green]123").
		SetTextColor(tview.Styles.SecondaryTextColor).
		SetAlign(tview.AlignRight).
		SetMaxWidth(30))

	table.SetBorder(true).SetTitle(" Connection ")

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(table, 0, 1, false)

	return flex
}
