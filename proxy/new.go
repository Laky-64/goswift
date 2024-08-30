package proxy

import (
	"github.com/blacktop/go-macho"
	"github.com/blacktop/go-macho/types"
	"reflect"
	"unsafe"
)

func New(file *macho.File) (*Context, error) {
	v := reflect.ValueOf(file).Elem().FieldByName("cr").Elem().Elem()
	return &Context{
		File: file,
		cr:   reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Interface().(*types.CustomSectionReader),
	}, nil
}
