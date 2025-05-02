package ui

import (
	"github.com/rivo/tview"
)

// CreateMenu creates the left menu.
func Menu() *tview.Flex {
	left := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green][1] SELECT[white]
[red][2] INSERT
[green][3] UPDATE
[red][4] DELETE
[green][5] CREATE
[red][6] DROP
[green][7] ALTER
[red][8] TRUNCATE

[red][9] WRITE ON LOG`)
	leftBox := tview.NewFlex().AddItem(left, 0, 1, false)
	leftBox.SetBorder(true).SetTitle(" Menu ")
	return leftBox
}
