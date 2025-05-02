package ui

import (
	"github.com/rivo/tview"
)

// CreateQueriesContent creates the content for the "Queries" section
func CreateQueriesContent() *tview.TextView {
	queries := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[blue]Query results or input go here...")
	queries.SetBorder(true).SetTitle(" Queries ")
	return queries
}
