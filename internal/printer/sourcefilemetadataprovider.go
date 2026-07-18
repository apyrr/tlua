package printer

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/tspath"
)

type SourceFileMetaDataProvider interface {
	GetSourceFileMetaData(path tspath.Path) *ast.SourceFileMetaData
}
