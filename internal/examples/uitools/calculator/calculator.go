package main

import (
	"os"

	"github.com/topofstack/cutego/widgets"

	"github.com/topofstack/cutego/internal/examples/uitools/calculator/ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	ui.NewCalculatorForm(nil).Show()

	widgets.QApplication_Exec()
}
