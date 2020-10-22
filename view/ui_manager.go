package view

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
)

type UIManager struct {
	uiRoot *UIRoot
}

var ins *UIManager

func UIManagerIns() *UIManager {
	if ins == nil {
		ins = &UIManager{}
	}
	return ins
}

func (u *UIManager) Init() {
	a := app.NewWithID("io.fyne.demo")
	a.SetIcon(theme.FyneLogo())

	w := a.NewWindow("Fyne Demo")

	u.uiRoot = NewUIRoot()

	w.SetContent(u.uiRoot.GetContent())

	w.ShowAndRun()
}
