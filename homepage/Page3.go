package homepage

import (
	"car_controller/config"
	"car_controller/util"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"time"
)

func Page3(w fyne.Window) fyne.CanvasObject {

	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}
	ctrlAltTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl | fyne.KeyModifierAlt}

	w.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("我们按下了Ctrl+Tab")
	})
	w.Canvas().AddShortcut(ctrlAltTab, func(shortcut fyne.Shortcut) {
		log.Println("我们按下了Ctrl+Alt+Tab")
	})

	stringBuilder := binding.NewString()

	sendControllerForward := func() {
		logStr := "button buttonForward!"
		err := util.SendInstructions("forward")
		if err != nil {
			logStr = err.Error()
		}
		util.LogSend(logStr)
	}
	sendControllerBackward := func() {
		logStr := "button buttonBackward!"
		err := util.SendInstructions("backward")
		if err != nil {
			logStr = err.Error()
		}
		util.LogSend(logStr)
	}
	sendControllerLeft := func() {
		logStr := "button buttonLeft!"
		err := util.SendInstructions("left")
		if err != nil {
			logStr = err.Error()
		}
		util.LogSend(logStr)
	}
	sendControllerRight := func() {
		logStr := "button buttonRight!"
		err := util.SendInstructions("right")
		if err != nil {
			logStr = err.Error()
		}
		util.LogSend(logStr)
	}

	sendControllerStop := func() {
		logStr := "button buttonStop!"
		err := util.SendInstructions("stop")
		if err != nil {
			logStr = err.Error()
		}
		util.LogSend(logStr)
	}

	buttonForward := widget.NewButtonWithIcon("前", theme.MoveUpIcon(), sendControllerForward)
	buttonBackward := widget.NewButtonWithIcon("后", theme.MoveDownIcon(), sendControllerBackward)
	buttonLeft := widget.NewButtonWithIcon("左", theme.NavigateBackIcon(), sendControllerLeft)
	buttonRight := widget.NewButtonWithIcon("右", theme.NavigateNextIcon(), sendControllerRight)
	buttonStop := widget.NewButtonWithIcon("Stop", theme.RadioButtonIcon(), sendControllerStop)

	controllerLogText := widget.NewMultiLineEntry()
	controllerLogText.Wrapping = fyne.TextWrapWord
	controllerLogText.Password = true
	controllerLogText.Resize(fyne.NewSize(100, 400))
	clock := widget.NewLabel("")
	go updateTime(clock)

	IPEntry := widget.NewEntryWithData(binding.BindString(&config.Config.Car.Ip))
	IPEntry.SetPlaceHolder("IP")

	portEntry := widget.NewEntryWithData(binding.BindString(&config.Config.Car.Port))
	portEntry.SetPlaceHolder("Port")
	portEntry.Validator = validation.NewRegexp(`\d`, "Must contain a number")
	//floats := container.NewGridWithColumns(2, IPEntry, portEntry)
	netButton := container.NewVScroll(container.NewVBox(
		clock,
		&widget.Separator{},
		IPEntry,
		portEntry,
		&widget.Separator{},
		&widget.Button{
			Text:       "创建连接",
			Importance: widget.HighImportance,
			Icon:       theme.ConfirmIcon(),
			OnTapped: func() {
				err := util.InitMessageConnect(config.Config.Car)
				if err != nil {
					util.LogSend(err.Error())
				} else {
					util.LogSend("已连接！")
				}
			},
		},
		&widget.Button{
			Text:       "测试连接",
			Importance: widget.SuccessImportance,
			Icon:       theme.MailSendIcon(),
			OnTapped: func() {
				err := util.TestInstructions()
				if err != nil {
					util.LogError(err)
				} else {
					util.LogSend("连接测试 success！")
				}
			},
		},
		&widget.Button{
			Text:       "断开连接",
			Importance: widget.DangerImportance,
			Icon:       theme.ContentClearIcon(),
			OnTapped: func() {
				err := util.UnMessageConnect()
				if err != nil {
					util.LogError(err)
				} else {
					util.LogSend("已断开连接！")
				}
			},
		},
	))

	operationLog := widget.NewEntryWithData(stringBuilder)
	operationLog.Disabled()       // 禁用输入，只用于显示
	operationLog.MultiLine = true // 允许多行输入
	go func() {
		for {
			select {
			case msg := <-util.LogManage.LogChan:
				currentLog, _ := stringBuilder.Get()
				timeNow := time.Now().Format("2006/1/02 15:04:05.000")
				newLog := "【" + timeNow + "】" + msg + "\n" + currentLog
				err := stringBuilder.Set(newLog)
				if err != nil {
					return
				}
			case <-util.LogManage.Done:
				return
			}
		}
	}()

	right := container.NewVSplit(
		netButton,
		operationLog,
	)
	blank := widget.NewLabel("")
	left1 := container.New(layout.NewGridLayout(3), blank, buttonForward, blank, buttonLeft, buttonStop, buttonRight, blank, buttonBackward, blank)

	clearLog := widget.NewButtonWithIcon("清空日志", theme.DeleteIcon(), func() {
		err := stringBuilder.Set("")
		if err != nil {
			return
		}
	})
	clearLog.Importance = widget.HighImportance
	left := container.NewBorder(nil, clearLog, nil, nil, left1)

	return container.NewHSplit(left, right)
}

func updateTime(clock *widget.Label) {
	for range time.Tick(time.Second) {
		timeNow := time.Now().Format("当前时间: 2006/1/02 15:04:05")
		clock.SetText(timeNow)
	}
}
