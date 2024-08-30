package proxy

import (
	"github.com/blacktop/go-macho"
	"github.com/blacktop/go-macho/types"
)

type Context struct {
	*macho.File
	cr *types.CustomSectionReader
}
