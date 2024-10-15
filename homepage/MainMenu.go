package homepage

import (
	"fyne.io/fyne/v2"
	"net/url"
)

func MakeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	connectCarItem := fyne.NewMenuItem("连接小车", nil)
	resetOrientationItem := fyne.NewMenuItem("重置舵机", nil)
	resetNetworkItem := fyne.NewMenuItem("切换网络", nil)
	controller := fyne.NewMenu("controller", connectCarItem, resetOrientationItem, resetNetworkItem)

	openUrl := func() {
		u, _ := url.Parse("https://byzhao.cn")
		_ = a.OpenURL(u)
	}
	helpMenu := fyne.NewMenu("help",
		fyne.NewMenuItem("Documentation", openUrl),
		fyne.NewMenuItem("Support", openUrl),
		fyne.NewMenuItemSeparator(), // 分割
		fyne.NewMenuItem("Sponsor", openUrl))

	main := fyne.NewMainMenu(controller, helpMenu)
	return main
}
