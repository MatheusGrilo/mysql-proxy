package ui

import (
	"github.com/rivo/tview"
)

// CreateCenter creates the central content area
func CreateLayout() *tview.Flex {
	menu := Menu()
	footer := CreateFooter()

	// Central content area
	connectionContent := CreateConnectionContent()
	queriesContent := CreateQueriesContent()

	center := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(connectionContent, 4, 0, false).
		AddItem(queriesContent, 0, 1, false)

		// Main layout
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(menu, 20, 1, false).
				AddItem(center, 0, 2, false),
			0, 1, false,
		).
		AddItem(footer, 3, 0, false) // Footer
	return flex
}
