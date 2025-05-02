package main

import (
	"github.com/matheusgrilo/mysql-proxy/ui"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Cria o layout principal usando o package ui
	layout := ui.CreateLayout()

	// Configura e executa o aplicativo
	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}
