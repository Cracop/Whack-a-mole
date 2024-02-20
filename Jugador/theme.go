package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct {
	fyne.Theme
}

func newMyTheme() fyne.Theme {
	return &myTheme{Theme: theme.DefaultTheme()}
}

func (m *myTheme) Color2(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m *myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == "azul" {

		return color.RGBA{254, 254, 254, 10}
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m *myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m *myTheme) Size(name fyne.ThemeSizeName) float32 {

	if name == theme.SizeNameText {
		return 12
	}

	return m.Theme.Size(name)
}
