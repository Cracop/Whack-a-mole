package main

import (
	"fyne.io/fyne/v2/app"
)

type CLIENT struct {
	loginGUI *LOGINGUI
	connect  *CONNECTION
	gameGUI  *GAMEGUI
}

func main() {
	// log.Println("Aqui llega")
	a := app.New()
	// w := a.NewWindow("VistaJugador")
	// client := CLIENT{}
	c := CONNECTION{}
	c.monster = make(chan int)
	c.start = false

	//TODO: hacer que el custom theme funcione
	// t := newMyTheme()
	// t.Color("azul", theme.VariantLight)
	// a.Settings().SetTheme(t)

	// w.SetMaster()

	// w.Resize(fyne.NewSize(400, 100))

	// gu := GUI{}

	// grid :=
	// makeGameGUI(a)

	makeLoginGUI(a, &c)

	a.Run()
}
