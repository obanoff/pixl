package swatch

import (
	"image/color"
	"pixl/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// simple square, possible to click on it to change the color
type Swatch struct {
	widget.BaseWidget
	Selected     bool            // whether or not the swatch has been selected
	Color        color.Color     // color of the swatch
	SwatchIndex  int             // used to select the swatch
	clickHandler func(s *Swatch) // functionality on click on the swatch
}

// set color of the swatch
func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh() // update the screen after color is changed
}

func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		clickHandler: clickHandler,
		SwatchIndex:  swatchIndex,
	}
	// to get default functionality
	swatch.ExtendBaseWidget(swatch)
	return swatch
}

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  swatch,
	}
}
