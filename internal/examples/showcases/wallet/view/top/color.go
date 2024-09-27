package top

import (
	"github.com/topofstack/cutego/quick"

	_ "github.com/topofstack/cutego/internal/examples/showcases/wallet/view/top/controller"
)

func init() { colorTemplate_QmlRegisterType2("TopTemplate", 1, 0, "ColorTemplate") }

type colorTemplate struct {
	quick.QQuickItem

	_ func() `signal:"change,->(controller.NewColorController(nil))"`
}
