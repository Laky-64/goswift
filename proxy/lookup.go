package proxy

import (
	"fmt"
	"github.com/Laky-64/goswift/demangling"
	"github.com/blacktop/go-macho/types/swift"
	"strings"
)

func (ctx *Context) lookup(kind demangling.SymbolicReferenceKind, addr uint64) (*demangling.Node, error) {
	var name string
	var nodeKind demangling.NodeKind
	var isType bool
	switch kind {
	case demangling.SymbolicContext:
		ptr, err := ctx.GetPointerAtAddress(addr)
		if err != nil {
			return nil, err
		}
		ptr = ctx.SlidePointer(ptr)
		if bind, err := ctx.GetBindName(ptr); err == nil {
			name = bind
			nodeKind = demangling.OpaqueTypeDescriptorSymbolicReferenceKind
		} else {
			if ptr == 0 {
				name, err = ctx.symbolLookup(addr)
				if err != nil {
					return nil, err
				}
				nodeKind = demangling.TypeSymbolicReferenceKind
				isType = true
			} else {
				if err = ctx.cr.SeekToAddr(ctx.SlidePointer(ptr)); err != nil {
					return nil, fmt.Errorf("failed to seek to indirect context descriptor: %v", err)
				}
				contextDesc, err := ctx.getContextDesc(ctx.SlidePointer(ptr))
				if err != nil {
					return nil, fmt.Errorf("failed to read indirect context descriptor: %v", err)
				}
				name = contextDesc.Name
				if len(contextDesc.Parent) > 0 {
					name = contextDesc.Parent + "." + name
				}
				if contextDesc.Flags.Kind() == swift.CDKindProtocol {
					nodeKind = demangling.ProtocolSymbolicReferenceKind
				} else if contextDesc.Flags.Kind() == swift.CDKindOpaqueType {
					nodeKind = demangling.OpaqueTypeDescriptorSymbolicReferenceKind
				} else {
					nodeKind = demangling.TypeSymbolicReferenceKind
					isType = true
				}
			}
		}
	default:
		return nil, nil
	}
	node := &demangling.Node{
		Kind: nodeKind,
		Text: name,
	}
	if context, err := demangling.New([]byte(node.Text)); err == nil {
		if nodeRes, err := context.Result(); err == nil {
			node = nodeRes
			for {
				switch node.Kind {
				case demangling.NominalTypeDescriptorKind,
					demangling.GlobalKind:
					node = node.FirstChild()
					continue
				default:
				}
				break
			}
		}
	}
	if node.Kind == demangling.TypeSymbolicReferenceKind {
		classInfo := strings.Split(node.Text, ".")
		if len(classInfo) > 1 {
			node = &demangling.Node{
				Kind: demangling.ClassKind,
				Children: []*demangling.Node{
					{
						Kind: demangling.ModuleKind,
						Text: strings.Join(classInfo[:len(classInfo)-1], "."),
					},
					{
						Kind: demangling.IdentifierKind,
						Text: classInfo[len(classInfo)-1],
					},
				},
			}
		}
	}
	if isType {
		node = demangling.CreateType(node)
	}
	return node, nil
}
