package demangling

func createStandardSubstitution(subSt rune, secondLevel bool) *Node {
	if std, found := standardTypes[subSt]; !secondLevel && found {
		return createSwiftType(std.Kind, std.TypeName)
	}
	if std, found := standardTypesConcurrency[subSt]; secondLevel && found {
		return createSwiftType(std.Kind, std.TypeName)
	}
	return nil
}
