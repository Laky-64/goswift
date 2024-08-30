package demangling

import (
	"bytes"
)

func prefixLength(mangled []byte) int {
	if len(mangled) == 0 {
		return 0
	}
	prefixes := []string{
		"_T0",       /*Swift 4*/
		"$S", "_$S", /*Swift 4.x*/
		"$s", "_$s", /*Swift 5*/
		"@__swiftmacro_", /*Swift 5+ for filenames*/
	}
	for _, prefix := range prefixes {
		if bytes.HasPrefix(mangled, []byte(prefix)) {
			return len(prefix)
		}
	}
	return 0
}
