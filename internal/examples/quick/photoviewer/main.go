//source: https://doc.qt.io/qt-5/qtquick-demos-photoviewer-example.html

package main

import (
	"os"

	"github.com/topofstack/cutego/core"
	"github.com/topofstack/cutego/gui"
	"github.com/topofstack/cutego/qml"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := gui.NewQGuiApplication(len(os.Args), os.Args)

	qtTranslator := core.NewQTranslator(nil)
	qtTranslator.Load("qml_"+core.QLocale_System().Name(), ":/i18n/", "", "")
	app.InstallTranslator(qtTranslator)

	engine := qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:///main.qml", 0))

	gui.QGuiApplication_Exec()
}
