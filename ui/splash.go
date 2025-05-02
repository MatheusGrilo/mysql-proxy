package ui

import (
	"fmt"
	"net"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/matheusgrilo/mysql-proxy/utils"
	"github.com/rivo/tview"
)

// CreateSplash creates a splash screen form and returns a boolean indicating whether to proceed.
// It also returns the MySQL IP, Port, and Proxy values entered by the user.
func CreateSplash(app *tview.Application) (bool, string, string, string) {
	var proceed bool
	var mysqlIP, mysqlPort, proxyPort string

	// Error text for IP
	mysqlIPError := tview.NewTextView().
		SetDynamicColors(true).
		SetText("").
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorRed)

	mysqlPortError := tview.NewTextView().
		SetDynamicColors(true).
		SetText("").
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorRed)

	proxyPortError := tview.NewTextView().
		SetDynamicColors(true).
		SetText("").
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorRed)

	// Default values
	mysqlIP = "172.20.80.253"
	mysqlPort = "3306"
	proxyPort = "4040"

	form := tview.NewForm().
		AddInputField("MySQL IP", mysqlIP, 20, nil, func(text string) {
			mysqlIP = text
		}).
		AddInputField("Port", mysqlPort, 20, nil, func(text string) {
			mysqlPort = text
		}).
		AddInputField("Proxy", proxyPort, 20, nil, func(text string) {
			proxyPort = text
		}).
		AddButton("Proceed", func() {
			// Validate MySQL IP
			if !utils.IsResolvable(mysqlIP) {
				mysqlIPError.SetText("[red]Invalid or empty IP address")
				return
			}

			// Validate MySQL Port
			port, err := strconv.Atoi(mysqlPort)
			if err != nil || port <= 0 || port > 65535 {
				mysqlPortError.SetText("[red]Invalid MySQL Port")
				return
			}

			// Check if we can telnet to MySQL
			conn, err := net.Dial("tcp", net.JoinHostPort(mysqlIP, mysqlPort))
			if err != nil {
				mysqlPortError.SetText("[red]Cannot connect to MySQL on this port")
				return
			}
			conn.Close()
			mysqlPortError.SetText("")

			// Validate Proxy Port
			proxyPortInt, err := strconv.Atoi(proxyPort)
			if err != nil || proxyPortInt <= 0 || proxyPortInt > 65535 {
				proxyPortError.SetText("[red]Invalid Proxy Port")
				return
			}

			// Check if Proxy Port is in use on localhost
			conn, err = net.Dial("tcp", fmt.Sprintf("127.0.0.1:%s", proxyPort))
			if err == nil {
				// If the connection succeeds, the port is in use
				conn.Close()
				proxyPortError.SetText("[red]Proxy Port is already in use on localhost")
				return
			}
			proxyPortError.SetText("")

			// If all validations pass
			mysqlIPError.SetText("")
			mysqlPortError.SetText("")
			proxyPortError.SetText("")
			proceed = true
			app.Stop()
		}).
		AddButton("Quit", func() {
			proceed = false
			app.Stop()
		})

	form.SetBorder(true).
		SetTitle(" Enter your MySQL Connection ").
		SetTitleAlign(tview.AlignLeft)

	// Layout vertical: formulário + erro
	formWithError := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true).
		AddItem(mysqlIPError, 1, 0, false).
		AddItem(mysqlPortError, 1, 0, false).
		AddItem(proxyPortError, 1, 0, false)

	// Layout centralizado na tela
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(formWithError, 0, 1, true)

	// Run
	if err := app.SetRoot(flex, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}

	return proceed, mysqlIP, mysqlPort, proxyPort
}
