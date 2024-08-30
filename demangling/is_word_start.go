package demangling

import "unicode"

func isWordStart(ch rune) bool {
	return !unicode.IsDigit(ch) && ch != '_' && ch != 0
}
