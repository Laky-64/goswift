package demangling

func getMangledLifetimeDependenceKind(nextChar rune) MangledLifetimeDependenceKind {
	switch nextChar {
	case 'i':
		return MangledLifetimeInherit
	case 's':
		return MangledLifetimeScope
	default:
		return UnknownLifetime
	}
}
