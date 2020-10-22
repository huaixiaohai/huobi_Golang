package view

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"github.com/huobirdcenter/huobi_golang/cmd/marketclientexample"
	"image/color"
	"time"
)

type UIRoot struct {
	canvasObject fyne.CanvasObject
}

func NewUIRoot() *UIRoot {
	root := &UIRoot{}
	root.makeContent()
	return root
}

func (u *UIRoot) GetContent() fyne.CanvasObject {
	return u.canvasObject
}

func (u *UIRoot) makeContent() {
	makeCell := func() fyne.CanvasObject {
		rect := canvas.NewRectangle(color.Black)
		rect.SetMinSize(fyne.NewSize(100, 100))

		text := canvas.NewText("222", color.White)

		textTime := canvas.NewText("333", color.White)

		go func() {
			for {
				text.Text = fmt.Sprintf("%f", marketclientexample.GetCandlestick("umausdt"))
				text.Refresh()
				textTime.Text = "2222" //time.Now().String() //fmt.Sprintf("%f", marketclientexample.GetCandlestick("forusdt"))
				textTime.Refresh()
				time.Sleep(time.Millisecond * 3)
			}
		}()

		con := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), canvas.NewText("1111", color.NRGBA{0xff, 0x00, 0x00, 0xff}), text, textTime)

		cell := fyne.NewContainerWithLayout(layout.NewCenterLayout(), rect, con)
		cell.Resize(fyne.NewSize(50, 50))
		return cell
	}

	u.canvasObject = fyne.NewContainerWithLayout(layout.NewGridLayout(1), makeCell())
}
