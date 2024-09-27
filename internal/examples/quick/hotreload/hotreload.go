package main

import (
	"os"
	"path/filepath"

	"github.com/topofstack/cutego/core"
	"github.com/topofstack/cutego/quick"
	"github.com/topofstack/cutego/widgets"
)

func initQQuickView(path string) *quick.QQuickView {

	var view = quick.NewQQuickView(nil)

	var watcher = core.NewQFileSystemWatcher2([]string{filepath.Dir(path)}, nil)

	var reload = func(p string) {
		println("changed:", p)
		view.SetSource(core.NewQUrl())
		view.Engine().ClearComponentCache()
		view.SetSource(core.NewQUrl3(path, 0))
	}

	//watcher.ConnectFileChanged(reload)
	watcher.ConnectDirectoryChanged(reload)

	return view
}

func main() {

	var path = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "bluszcz", "qt", "internal", "examples", "quick", "hotreload", "qml", "hotreload.qml")

	widgets.NewQApplication(len(os.Args), os.Args)

	var view = initQQuickView(path)
	view.SetSource(core.NewQUrl3(path, 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	widgets.QApplication_Exec()
}
