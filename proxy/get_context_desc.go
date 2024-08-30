package proxy

import (
	"fmt"
	"github.com/blacktop/go-macho/types/swift"
)

func (ctx *Context) getContextDesc(addr uint64) (context *swift.TargetModuleContext, err error) {
	var ptr uint64

	if (addr & 1) == 1 {
		addr = addr &^ 1
		ptr, err = ctx.GetPointerAtAddress(addr)
		if err != nil {
			return nil, fmt.Errorf("failed to read swift context descriptor pointer at address %#x: %v", addr, err)
		}
		ptr = ctx.SlidePointer(ptr)
	} else {
		ptr = addr
	}
	if err = ctx.cr.SeekToAddr(ptr); err != nil {
		if bind, err := ctx.GetBindName(ptr); err == nil {
			return &swift.TargetModuleContext{Name: bind}, nil
		} else if symbols, err := ctx.FindAddressSymbols(ptr); err == nil {
			if len(symbols) > 0 {
				for _, s := range symbols {
					if !s.Type.IsDebugSym() {
						return &swift.TargetModuleContext{Name: s.Name}, nil
					}
				}
			}
		}
	}
	context = &swift.TargetModuleContext{}
	if err = context.TargetModuleContextDescriptor.Read(ctx.cr, ptr); err != nil {
		return nil, fmt.Errorf("failed to read swift context descriptor: %w", err)
	}
	if context.ParentOffset.IsSet() {
		parent, err := ctx.getContextDesc(context.ParentOffset.GetAddress())
		if err != nil {
			return nil, fmt.Errorf("failed to read swift context descriptor parent context: %w", err)
		}
		if parent.Parent != "" {
			if parent.Name != "" {
				context.Parent = parent.Parent + "." + parent.Name
			} else {
				context.Parent = parent.Parent
			}
		} else {
			context.Parent = parent.Name
		}
	}

	switch context.Flags.Kind() {
	case swift.CDKindModule, swift.CDKindProtocol, swift.CDKindClass, swift.CDKindStruct, swift.CDKindEnum:
		if context.NameOffset.IsSet() {
			context.Name, err = ctx.GetCString(context.NameOffset.GetAddress())
			if err != nil {
				return nil, fmt.Errorf("failed to read swift module context name: %w", err)
			}
		}
	}
	return context, nil
}
