//go:build !qml
// +build !qml

package artist

import (
	"github.com/topofstack/cutego/core"
	"github.com/topofstack/cutego/widgets"

	"github.com/topofstack/cutego/internal/examples/sql/masterdetail_qml/controller"
)

type artistController struct {
	widgets.QGroupBox

	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"listModel"`

	//->controller
	_ func(row int) `signal:"changeArtist"`

	artistView *widgets.QComboBox
}

func (a *artistController) init() {

	a.artistView = widgets.NewQComboBox(nil)
	a.artistView.SetModel(controller.Instance.ArtistModel())

	layout := widgets.NewQGridLayout2()
	layout.AddWidget2(a.artistView, 0, 0, 0)
	a.SetLayout(layout)

	//

	//->controller
	a.artistView.ConnectCurrentIndexChanged(controller.Instance.ChangeArtist)
}
