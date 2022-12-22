package pxcanvas

import (
	"pixl/pxcanvas/brush"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

// Scrollable interface implementation
func (pxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

// Hoverable interface implementation
func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	if x, y := pxCanvas.MouseToCanvasXY(ev); x != nil && y != nil {
		brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
		cursor := brush.Cursor(pxCanvas.PxCanvasConfig, pxCanvas.appState.BrushType, ev, *x, *y)
		pxCanvas.renderer.SetCursor(cursor)
	} else {
		pxCanvas.renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh()
	pxCanvas.mouseState.previousCoord = &ev.PointEvent
}

// Hoverable interface implementation
func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseOut()                      {}

// Mousable interface implementation
func (pxCanvas *PxCanvas) MouseUp(ev *desktop.MouseEvent) {}

func (pxCanvas *PxCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
}
