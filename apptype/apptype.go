package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

type BrushType = int

// fyne config
type PxCanvasConfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSize         int
}

// for general application state
type State struct {
	BrushColor color.Color
	BrushType  int
	// Swatch will be some sort of a slice and SwatchSelected is just an index in that slice
	SwatchSelected int
	// path for current file on disk
	FilePath string
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}

type Brushable interface {
	SetColor(c color.Color, x, y int)
	MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int)
}
