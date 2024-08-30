package proxy

import (
	"fmt"
)

func (ctx *Context) symbolLookup(addr uint64) (string, error) {
	var err error
	var ptr uint64

	if (addr & 1) == 1 {
		addr = addr &^ 1
		ptr, err = ctx.GetPointerAtAddress(addr)
		if err != nil {
			return "", fmt.Errorf("failed to read protocol pointer @ %#x: %v", addr, err)
		}
		ptr = ctx.SlidePointer(ptr)
	} else {
		ptr = addr
	}
	if bind, err := ctx.GetBindName(ptr); err == nil {
		return bind, nil
	} else if symbols, err := ctx.FindAddressSymbols(ptr); err == nil {
		if len(symbols) > 0 {
			for _, s := range symbols {
				if !s.Type.IsDebugSym() {
					return s.Name, nil
				}
			}
		}
	}
	return "", fmt.Errorf("failed to find symbol for address %#x", addr)
}
