//source: https://github.com/StratifyLabs/StratifyQML

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
	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	widgets.QApplication_Exec()
}
