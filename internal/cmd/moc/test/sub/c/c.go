package c

import (
	"github.com/topofstack/cutego/core"

	_ "github.com/topofstack/cutego/internal/cmd/moc/test/sub/conf"
)

type StructSubGoC struct{}
type StructSubMocC struct{ core.QObject }
