//source: https://doc.qt.io/qt-5/qtwebview-minibrowser-example.html

package main

import (
	"os"

	"github.com/topofstack/cutego/core"
	"github.com/topofstack/cutego/gui"
	"github.com/topofstack/cutego/qml"
	"github.com/topofstack/cutego/webview"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)
	webview.QtWebView_Initialize()

	var view = qml.NewQQmlApplicationEngine(nil)
	view.Load(core.NewQUrl3("qrc:/main.qml", 0))

	gui.QGuiApplication_Exec()
}
