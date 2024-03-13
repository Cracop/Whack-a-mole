package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type LOGINGUI struct {
	layout fyne.CanvasObject
	app    fyne.App
}

func makeLoginGUI(app fyne.App, c *CONNECTION) {
	g := LOGINGUI{}
	g.app = app

	w := app.NewWindow("Login")

	ipEntry := widget.NewEntry()
	ipEntry.SetPlaceHolder("Enter IP Address")
	nombreEntry := widget.NewEntry()
	puertoEntry := widget.NewEntry()
	nombreEntry.SetPlaceHolder("Nombre jugador")
	puertoEntry.SetPlaceHolder("Puerto")

	ipEntry.Text = c.ipAddress
	puertoEntry.Text = c.puerto
	nombreEntry.Text = c.nombre

	joinButton := widget.NewButton("Join", func() {
		// ip := ipEntry.Text
		// Do something with the entered IP address
		c.ipAddress = ipEntry.Text
		c.puerto = puertoEntry.Text
		c.nombre = nombreEntry.Text

		makeGameGUI(app, c)
		if c.conn == nil {
			LoginTCP(c)
			go receiveUDP(c)
		}

		w.Close()
	})

	g.layout = container.New(layout.NewGridLayout(3),
		nombreEntry,
		ipEntry,
		puertoEntry,
		layout.NewSpacer(),
		joinButton,
		layout.NewSpacer(),
	)

	w.SetContent(g.layout)
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(500, 100))
	w.Show()
}
