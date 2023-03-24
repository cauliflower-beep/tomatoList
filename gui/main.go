package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"os"
)

func init() {
	_ = os.Setenv("FYNE_FONT", "./anna.ttf")
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("tomatoList")
	myWindow.Resize(fyne.NewSize(400, 600))

	entry := widget.NewEntry()
	entry.SetPlaceHolder("请输入待办事项...")
	entry.Resize(fyne.NewSize(3000, entry.MinSize().Height))
	myWindow.SetContent(container.NewVBox(
		widget.NewLabelWithStyle("番茄清单", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Monospace: true, Symbol: true}),
		container.NewHBox(
			entry,
			widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {}),
		),
	))

	myWindow.ShowAndRun()

	os.Unsetenv("FYNE_FONT")
}
