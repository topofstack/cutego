package controller

import (
	"github.com/topofstack/cutego/core"

	_ "github.com/topofstack/cutego/internal/examples/showcases/wallet/files/dialog/controller"
)

var ActionButtonController *actionButtonController

type actionButtonController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) `signal:"showDownload,->(controller.Controller)"`
	_ func(string) `signal:"deleteRequest,auto"`
}

func (c *actionButtonController) init() {
	ActionButtonController = c
}

func (c *actionButtonController) deleteRequest(name string) {
}
