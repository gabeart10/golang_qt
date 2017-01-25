package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

func CreateBullet() (Bullet *widgets.QGraphicsRectItem) {
	Bullet = widgets.NewQGraphicsRectItem(nil)
	Bullet.SetRect2(0, 0, 10, 50)
	var time = core.NewQTimer(nil)
	time.ConnectTimeout(func() {
		Bullet.SetPos2(Bullet.X(), Bullet.Y()-10)
	})
	time.Start(50)
	return
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var scene = widgets.NewQGraphicsScene(nil)

	var rect = widgets.NewQGraphicsRectItem(nil)
	rect.SetRect2(0, 0, 100, 100)
	rect.SetFlag(widgets.QGraphicsItem__ItemIsFocusable, true)
	rect.SetFocus(core.Qt__ActiveWindowFocusReason)

	rect.ConnectKeyPressEvent(func(keypress *gui.QKeyEvent) {
		if keypress.Key() == int(core.Qt__Key_Left) {
			rect.SetPos2(rect.X()-10, rect.Y())
		} else if keypress.Key() == int(core.Qt__Key_Right) {
			rect.SetPos2(rect.X()+10, rect.Y())
		} else if keypress.Key() == int(core.Qt__Key_Down) {
			rect.SetPos2(rect.X(), rect.Y()+10)
		} else if keypress.Key() == int(core.Qt__Key_Up) {
			rect.SetPos2(rect.X(), rect.Y()-10)
		} else if keypress.Key() == int(core.Qt__Key_Space) {
			var bullet = CreateBullet()
			bullet.SetPos2(rect.X(), rect.Y())
			scene.AddItem(bullet)
		}
	})

	scene.AddItem(rect)

	var view = widgets.NewQGraphicsView2(scene, nil)
	view.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	view.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)

	view.Show()

	widgets.QApplication_Exec()
}
