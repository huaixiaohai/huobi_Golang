package widget

import (
	"fyne.io/fyne/widget"
)

type SymbolCell struct {
	widget.BaseWidget
	widget.Button
}

func NewSymbolCell() *SymbolCell {
	return &SymbolCell{}
}

