package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
	"strconv"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	var window = widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(400, 400)
	window.SetWindowTitle("gui")
	var centralWidget = widgets.NewQWidget(window, 0)
	window.SetCentralWidget(centralWidget)

	var vLayout = widgets.NewQVBoxLayout2(centralWidget)
	var hLayout = widgets.NewQHBoxLayout()

	var gridLayout = widgets.NewQGridLayout2()
	gridLayout.SetSpacing(0)

	hLayout.AddStretch(1)
	hLayout.AddLayout(gridLayout, 1)

	vLayout.AddLayout(hLayout, 1)

	var buttons = make([]*widgets.QPushButton, 100)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			var button = widgets.NewQPushButton2(strconv.Itoa(i*10+j), nil)
			button.SetFixedSize2(40, 40)
			button.ConnectClicked(func(flag bool) {
				widgets.QMessageBox_Information(nil, "OK", strconv.Itoa(i*10)+"+"+strconv.Itoa(j)+"="+strconv.Itoa(i*10+j), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			})
			buttons = append(buttons, button)
			gridLayout.AddWidget(button, i, j, 0)
			gridLayout.SetColumnMinimumWidth(j, 40)
		}
		gridLayout.SetRowMinimumHeight(i, 40)
	}

	for i := 0; i < 100; i++ {
	}
	window.Show()

	widgets.QApplication_Exec()

}
