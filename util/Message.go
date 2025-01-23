package util

import "fyne.io/fyne/v2"

func SendMessage(title string, content string) {
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   title,
		Content: content,
	})
}
