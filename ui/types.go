package ui

import (
	"awesomeProject/apptype"
	"awesomeProject/pxcanvas"
	"awesomeProject/swatch"
	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
