//source: http://blog.lasconic.com/share-on-ios-and-android-using-qml/

package main

import (
	"os"

	"github.com/topofstack/cutego/core"
	"github.com/topofstack/cutego/qml"
	"github.com/topofstack/cutego/widgets"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	widgets.NewQApplication(len(os.Args), os.Args)

	engine := qml.NewQQmlApplicationEngine(nil)
	engine.RootContext().SetContextProperty("shareUtils", NewShareUtils(nil))
	engine.Load(core.NewQUrl3("qrc:/main.qml", 0))

	widgets.QApplication_Exec()
}
