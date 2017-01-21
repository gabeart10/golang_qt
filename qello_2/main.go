package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(200, 200)
	window.SetWindowTitle("Qello_2")

	var layout = widgets.NewQVBoxLayout()

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(layout)

	var input = widgets.NewQTextEdit(nil)

	var button = widgets.NewQPushButton2("Ok", nil)
	button.ConnectClicked(func(flag bool) {
		widgets.QMessageBox_Information(nil, "Ok", "Hello, "+input.ToPlainText(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		input.Clear()
	})

	input.SetMaximumHeight(20)
	input.SetMaximumWidth(100)

	layout.AddWidget(input, 0, core.Qt__AlignCenter)
	layout.AddWidget(button, 0, core.Qt__AlignCenter)

	window.SetCentralWidget(centralWidget)
	window.Show()

	widgets.QApplication_Exec()
}
