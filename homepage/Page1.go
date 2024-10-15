package homepage

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MakeHomePage(myApp fyne.App, w fyne.Window) fyne.CanvasObject {

	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			myApp.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})
		}),
		widget.NewButton("Light", func() {
			myApp.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})
		}),
	)
	return container.NewBorder(nil, themes, nil, nil, nil)
}
