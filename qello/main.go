package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Qello")
	window.SetMinimumSize2(200, 200)

	var layout = widgets.NewQVBoxLayout()

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(layout)

	var input = widgets.NewQTextEdit(nil)

	var text = widgets.NewQLabel(nil, 0)

	var font = gui.NewQFont()
	font.SetFamily("Georgia")
	font.SetPointSize(20)

	text.SetFont(font)

	input.ConnectTextChanged(func() {
		text.SetText("Hi, " + input.ToPlainText())
	})

	layout.AddWidget(input, 0, core.Qt__AlignTop)
	layout.AddWidget(text, 0, core.Qt__AlignBottom)

	window.SetCentralWidget(centralWidget)
	window.Show()

	widgets.QApplication_Exec()

}
