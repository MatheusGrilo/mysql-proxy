package main

import (
	"github.com/matheusgrilo/mysql-proxy/ui"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Exibe a tela de splash
	proceed, mysqlIP, mysqlPort, proxyPort := ui.CreateSplash(app)
	if proceed {
		// Se o usuário clicar em "Proceed", executa o layout principal
		layout := ui.CreateLayout(mysqlIP, mysqlPort, proxyPort)
		if err := app.SetRoot(layout, true).Run(); err != nil {
			panic(err)
		}
	} else {
		// Se o usuário clicar em "Quit", encerra o programa
		app.Stop()
	}
}
