package main

import (
	"github.com/topofstack/cutego/core"
)

type Person struct {
	core.QObject

	_ string `property:"name"`
	_ int    `property:"shoeSize"`
}
