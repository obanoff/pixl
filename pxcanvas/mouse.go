package pxcanvas

import (
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
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh()
	pxCanvas.mouseState.previousCoord = &ev.PointEvent
}

// Hoverable interface implementation
func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {
	// pxCanvas.mouseState.hovering = true
}
func (pxCanvas *PxCanvas) MouseOut() {
	// pxCanvas.mouseState.hovering = false
}
