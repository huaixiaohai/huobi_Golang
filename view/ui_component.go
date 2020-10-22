package view

import "fyne.io/fyne"

type UIComponent struct {
	canvasObject *fyne.CanvasObject
}

func NewUIComponent(canvasObj *fyne.CanvasObject) *UIComponent {
	return &UIComponent{
		canvasObject: canvasObj,
	}
}

func (u *UIComponent) GetContent() *fyne.CanvasObject {
	return u.canvasObject
}
