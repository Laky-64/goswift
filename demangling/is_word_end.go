package demangling

import "unicode"

func isWordEnd(ch, prevCh rune) bool {
	if ch == '_' || ch == 0 {
		return true
	}
	if !unicode.IsUpper(prevCh) && unicode.IsUpper(ch) {
		return true
	}
	return false
}
