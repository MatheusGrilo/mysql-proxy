package ui

import (
	"github.com/rivo/tview"
)

// CreateFooter creates the footer.
func CreateFooter() *tview.TextView {
	footer := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("[blue]MatheusGrilo[default] | [green]GitHub.com/MatheusGrilo/mysql-proxy[default] | [red]v0.1")
		// SetText("A [red]c[yellow]o[green]l[darkcyan]o[blue]r[darkmagenta]f[red]u[yellow]l[white] [black:red]c[:yellow]o[:green]l[:darkcyan]o[:blue]r[:darkmagenta]f[:red]u[:yellow]l[white:-] [::bu]title")
	footer.SetBorder(true)
	return footer
}
