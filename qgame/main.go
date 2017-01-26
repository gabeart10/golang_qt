package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

var scene = widgets.NewQGraphicsScene(nil)

func CreateBullet() (Bullet *widgets.QGraphicsRectItem) {
	Bullet = widgets.NewQGraphicsRectItem(nil)
	Bullet.SetRect2(0, 0, 10, 50)
	var time = core.NewQTimer(nil)
	time.ConnectTimeout(func() {
		Bullet.SetPos2(Bullet.X(), Bullet.Y()-10)
		if Bullet.Y() < -50 {
			scene.RemoveItem(Bullet)
		}
	})
	time.Start(50)
	return
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var player = widgets.NewQGraphicsRectItem(nil)
	player.SetRect2(0, 0, 100, 100)
	player.SetFlag(widgets.QGraphicsItem__ItemIsFocusable, true)
	player.SetFocus(core.Qt__ActiveWindowFocusReason)

	player.ConnectKeyPressEvent(func(keypress *gui.QKeyEvent) {
		if keypress.Key() == int(core.Qt__Key_Left) {
			player.SetPos2(player.X()-10, player.Y())
		} else if keypress.Key() == int(core.Qt__Key_Right) {
			player.SetPos2(player.X()+10, player.Y())
		} else if keypress.Key() == int(core.Qt__Key_Space) {
			var bullet = CreateBullet()
			scene.AddItem(bullet)
			bullet.SetPos2(player.X(), player.Y())
		}
	})

	scene.AddItem(player)

	var view = widgets.NewQGraphicsView2(scene, nil)
	view.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	view.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)

	view.Show()
	view.SetFixedSize2(800, 600)
	scene.SetSceneRect2(0, 0, 800, 600)

	player.SetPos2(float64(view.Width())/2, float64(view.Height())-player.Rect().Height())

	widgets.QApplication_Exec()
}
