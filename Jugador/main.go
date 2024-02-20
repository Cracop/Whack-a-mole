package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {

	a := app.New()
	// w := a.NewWindow("VistaJugador")
	c := CONNECTION{}

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
