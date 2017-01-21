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
	window.SetWindowTitle("Fonts")
	window.SetMinimumSize2(200, 200)

	var layout = widgets.NewQVBoxLayout()

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(layout)

	var text = widgets.NewQLabel(nil, 0)

	var font = gui.NewQFont()

	font.SetFamily("Georgia")

	font.SetPointSize(40)

	text.SetText("Test of the wonderful fonts 😍 ")
	text.SetFont(font)

	layout.AddWidget(text, 0, core.Qt__AlignCenter)

	window.SetCentralWidget(centralWidget)
	window.Show()

	widgets.QApplication_Exec()

}
