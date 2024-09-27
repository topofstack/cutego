//source: https://github.com/lirios/fluid

package main

import (
	"os"

	"github.com/topofstack/cutego/core"
	"github.com/topofstack/cutego/gui"
	"github.com/topofstack/cutego/qml"
	"github.com/topofstack/cutego/quickcontrols2"
)

func main() {

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	app.SetQuitOnLastWindowClosed(true)

	if quickcontrols2.QQuickStyle_Name() == "" {
		quickcontrols2.QQuickStyle_SetStyle("Material")
	}

	engine := qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	app.Exec()
}
