package main

import (
	"car_controller/homepage"
	"car_controller/util"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.New()
	mainWindow := myApp.NewWindow("小车控制器")

	initWindowSetting(myApp, mainWindow)
	mainWindow.SetMainMenu(homepage.MakeMenu(myApp, mainWindow)) // 设置菜单
	mainWindow.SetMaster()

	//mainWindow.SetContent(clock)

	//grid := container.New(layout.NewGridLayout(3), text1, text2, text3, image)
	//mainWindow.SetContent(grid)

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("首页1", theme.HomeIcon(), homepage.MakeHomePage(myApp, mainWindow)), // 首页
		container.NewTabItemWithIcon("测试2", theme.ComputerIcon(), homepage.Page2(myApp, mainWindow)),
		container.NewTabItemWithIcon("测试3", theme.ContentPasteIcon(), homepage.Page3(mainWindow)),
		container.NewTabItemWithIcon("测试4", theme.MailSendIcon(), homepage.Page4(mainWindow)),
		container.NewTabItem("标签 1", widget.NewLabel("你好")),
	)
	tabs.Append(container.NewTabItem("标签 2", widget.NewLabel("你好")))
	tabs.SetTabLocation(container.TabLocationBottom)
	mainWindow.SetContent(tabs)
	//text1 := canvas.NewText("你好", color.Black)
	//text2 := canvas.NewText("在那里", color.Black)
	//text3 := canvas.NewText("(右侧)", color.Black)
	//content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)
	//
	//text4 := canvas.NewText("居中", color.Black)
	//centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	//mainWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))

	mainWindow.Show()
	myApp.Run()

	tidyUp()
}

func tidyUp() {
	err := util.UnMessageConnect()
	if err == nil {
		util.LogSend("断开socket连接")
	}
	fmt.Println("控制器程序退出")
}

// 初始化master窗口
func initWindowSetting(myApp fyne.App, mainWindow fyne.Window) {
	iconResource, err := fyne.LoadResourceFromPath("Icon.png")
	if err != nil {
		fyne.LogError("Could not load icon", err)
	} else {
		mainWindow.SetIcon(iconResource)
	}
	myApp.Settings().SetTheme(theme.DefaultTheme())
	//mainWindow.Resize(fyne.NewSize(800, 600))
}
