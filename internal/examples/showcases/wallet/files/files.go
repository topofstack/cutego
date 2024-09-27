package files

import (
	"github.com/topofstack/cutego/core"
	"github.com/topofstack/cutego/quick"

	"github.com/topofstack/cutego/internal/examples/showcases/wallet/files/controller"
	"github.com/topofstack/cutego/internal/examples/showcases/wallet/files/dialog"
)

func init() { filesTemplate_QmlRegisterType2("FilesTemplate", 1, 0, "FilesTemplate") }

type filesTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"FilesModel"`
}

func (t *filesTemplate) init() {
	c := controller.NewFilesController(nil)

	t.SetFilesModel(c.Model().Filter)

	//needed here, because those are non qml views
	dialog.NewFilesUploadTemplate(nil)
	dialog.NewFolderUploadTemplate(nil)
}
