package demangling

import (
	"bytes"
	"fmt"
)

func New(mangled []byte) (*Context, error) {
	prefixLen := prefixLength(mangled)
	if bytes.HasPrefix(mangled, []byte("_T")) {
		return nil, fmt.Errorf("swift 4 mangled names are not supported")
	}
	return &Context{
		Data: mangled,
		Pos:  prefixLen,
		Size: len(mangled),
	}, nil
}
