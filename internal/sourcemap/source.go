package sourcemap

import "github.com/apyrr/tlua/internal/core"

type Source interface {
	Text() string
	FileName() string
	ECMALineMap() []core.TextPos
}
