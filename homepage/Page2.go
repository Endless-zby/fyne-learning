package homepage

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
)

func Page2(myApp fyne.App, w fyne.Window) fyne.CanvasObject {

	text1 := canvas.NewText("区域1", color.Black)
	image := canvas.NewImageFromFile("Icon.png")
	text2 := canvas.NewText("区域2", color.Black)
	text3 := canvas.NewText("区域3", color.Black)
	//grid := container.New(layout.NewGridLayout(3), text1, text2, text3, image)
	//mainWindow.SetContent(grid)

	return container.New(layout.NewGridLayout(3), text1, text2, text3, image)
}
