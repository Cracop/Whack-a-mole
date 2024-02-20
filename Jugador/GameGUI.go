package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type GAMEGUI struct {
	layout  fyne.CanvasObject
	cells   [9]*widget.Button // Store pointers to buttons
	labels  [3]canvas.Text
	info    [3]canvas.Text
	app     fyne.App
	connect *CONNECTION
}

func (g *GAMEGUI) pressButton(index int, c *CONNECTION) {
	// Access the value of the clicked cell
	log.Println("Cell tapped:", index+1)
	log.Println(g.connect.ipAddress)
	indexstr := strconv.Itoa(index)
	whackTCP(c, indexstr)
	//Maybe pasar esto a una subrutina
	g.cells[index].Importance = widget.HighImportance
	g.cells[index].Refresh()
	time.Sleep(100 * time.Millisecond)
	g.cells[index].Importance = widget.MediumImportance
	g.cells[index].Refresh()

}

func (g *GAMEGUI) buildButtons(c *CONNECTION) {
	for i := 0; i < len(g.cells); i++ {
		// Create button with label indicating cell number
		g.cells[i] = widget.NewButton(fmt.Sprintf("Cell %d", i+1), func(index int) func() {
			return func() {
				g.pressButton(index, c)
			}
		}(i))
	}
}

func /*(g *GUI)*/ makeGameGUI(app fyne.App, c *CONNECTION) /*GUI*/ {
	g := GAMEGUI{}
	g.app = app
	g.connect = c
	w := app.NewWindow("Client")
	g.buildButtons(c) // Build buttons
	exitBTN := widget.NewButton("Exit", func() {
		makeLoginGUI(app, c)
		w.Close()
	})

	// Create grid layout with buttons
	g.layout = container.New(layout.NewGridLayout(3),
		g.cells[0], g.cells[1], g.cells[2],
		g.cells[3], g.cells[4], g.cells[5],
		g.cells[6], g.cells[7], g.cells[8],
		layout.NewSpacer(), exitBTN, layout.NewSpacer(),
	)

	w.SetContent(g.layout)
	w.Resize(fyne.NewSize(300, 100))

	go play(app, c, &g)

	w.Show()

}

func play(app fyne.App, c *CONNECTION, g *GAMEGUI) {

}
