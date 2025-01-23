package homepage

import (
	"car_controller/util"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MakeHomePage(myApp fyne.App, w fyne.Window) fyne.CanvasObject {

	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			myApp.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})
			util.SendMessage("主题切换", "Dark....")
		}),
		widget.NewButton("Light", func() {
			myApp.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})
			util.SendMessage("主题切换", "Light....")
		}),
	)
	return container.NewBorder(nil, themes, nil, nil, nil)
}
