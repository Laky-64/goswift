package demangling

import "fmt"

// Reference:
// https://github.com/swiftlang/swift/blob/4987c3b970036046ba668b0fe779297e37fa9544/lib/Demangling/Demangler.cpp#L2400
func (ctx *Context) metaType() (*Node, error) {
	switch ctx.nextChar() {
	case 'a':
		return ctx.createWithPoppedType(TypeMetadataAccessFunctionKind), nil
	case 'A':
		return createWithChildren(ReflectionMetadataAssocTypeDescriptorKind, ctx.popProtocolConformance()), nil
	case 'b':
		return ctx.createWithPoppedType(CanonicalSpecializedGenericTypeMetadataAccessFunctionKind), nil
	case 'B':
		return createWithChildren(ReflectionMetadataBuiltinDescriptorKind, ctx.popNodeKind(TypeKind)), nil
	case 'c':
		return createWithChildren(ProtocolConformanceDescriptorKind, ctx.popProtocolConformance()), nil
	case 'C':
		tY := ctx.popNodeKind(TypeKind)
		if tY == nil || !isAnyGeneric(tY.FirstChild().Kind) {
			return nil, fmt.Errorf("expected a generic type")
		}
		return createWithChildren(ReflectionMetadataSuperclassDescriptorKind, tY.FirstChild()), nil
	case 'D':
		return ctx.createWithPoppedType(TypeMetadataDemanglingCacheKind), nil
	case 'f':
		return ctx.createWithPoppedType(FullTypeMetadataKind), nil
	case 'F':
		return createWithChildren(ReflectionMetadataFieldDescriptorKind, ctx.popNodeKind(TypeKind)), nil
	case 'g':
		return createWithChildren(OpaqueTypeDescriptorAccessorKind, ctx.popNode()), nil
	case 'h':
		return createWithChildren(OpaqueTypeDescriptorAccessorImplKind, ctx.popNode()), nil
	case 'i':
		return ctx.createWithPoppedType(TypeMetadataInstantiationFunctionKind), nil
	case 'I':
		return ctx.createWithPoppedType(TypeMetadataInstantiationCacheKind), nil
	case 'j':
		return createWithChildren(OpaqueTypeDescriptorAccessorKeyKind, ctx.popNode()), nil
	case 'J':
		return createWithChildren(NoncanonicalSpecializedGenericTypeMetadataCacheKind, ctx.popNode()), nil
	case 'k':
		return createWithChildren(OpaqueTypeDescriptorAccessorVarKind, ctx.popNode()), nil
	case 'K':
		return createWithChildren(MetadataInstantiationCacheKind, ctx.popNode()), nil
	case 'l':
		return ctx.createWithPoppedType(TypeMetadataSingletonInitializationCacheKind), nil
	case 'L':
		return createWithChildren(TypeMetadataLazyCacheKind, ctx.popNode()), nil
	case 'm':
		return ctx.createWithPoppedType(MetaclassKind), nil
	case 'M':
		return createWithChildren(CanonicalSpecializedGenericMetaclassKind, ctx.popNode()), nil
	case 'n':
		return ctx.createWithPoppedType(NominalTypeDescriptorKind), nil
	case 'N':
		return createWithChildren(NoncanonicalSpecializedGenericTypeMetadataKind, ctx.popNode()), nil
	case 'o':
		return ctx.createWithPoppedType(ClassMetadataBaseOffsetKind), nil
	case 'p':
		return createWithChildren(ProtocolDescriptorKind, ctx.popProtocol()), nil
	case 'P':
		return ctx.createWithPoppedType(GenericTypeMetadataPatternKind), nil
	case 'q':
		return createWithChildren(UniquableKind, ctx.popNode()), nil
	case 'Q':
		return createWithChildren(OpaqueTypeDescriptorKind, ctx.popNode()), nil
	case 'r':
		return ctx.createWithPoppedType(TypeMetadataCompletionFunctionKind), nil
	case 's':
		return ctx.createWithPoppedType(ObjCResilientClassStubKind), nil
	case 'S':
		return createWithChildren(ProtocolSelfConformanceDescriptorKind, ctx.popProtocol()), nil
	case 't':
		return createWithChildren(FullObjCResilientClassStubKind, ctx.popProtocol()), nil
	case 'u':
		return ctx.createWithPoppedType(MethodLookupFunctionKind), nil
	case 'U':
		return ctx.createWithPoppedType(ObjCMetadataUpdateFunctionKind), nil
	case 'V':
		return createWithChildren(PropertyDescriptorKind, ctx.popNodePred(isEntity)), nil
	case 'X':
		return ctx.privateContextDescriptor()
	case 'z':
		return ctx.createWithPoppedType(CanonicalPrespecializedGenericTypeCachingOnceTokenKind), nil
	default:
		return nil, fmt.Errorf("unexpected metatype kind: %c", ctx.Data[ctx.Pos-1])
	}
}
