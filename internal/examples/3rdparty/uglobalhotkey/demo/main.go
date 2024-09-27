package main

import (
	"os"

	"github.com/topofstack/cutego/widgets"

	"github.com/topofstack/cutego/internal/examples/3rdparty/uglobalhotkey/UGlobalHotkey"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	hk := UGlobalHotkey.NewUGlobalHotkeys(nil)
	hk.RegisterHotkey("Ctrl+Shift+A", 1)
	hk.ConnectActivated(func(id uint) {
		println("Activated:", id)
	})

	hk = UGlobalHotkey.NewUGlobalHotkeys(nil)
	hk.RegisterHotkey("Ctrl+Shift+B", 2)
	hk.ConnectActivated(func(id uint) {
		println("Activated:", id)
	})

	widgets.QApplication_Exec()
}
