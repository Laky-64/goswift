package demangling

import (
	"fmt"
	"strings"
	"unicode"
)

func (ctx *Context) identifier() (*Node, error) {
	var hasWordSubSts bool
	var isPunyCoded bool
	c := ctx.peekChar()
	if !unicode.IsDigit(c) {
		return nil, fmt.Errorf("expected digit, got %c", c)
	}
	if c == '0' {
		ctx.nextChar()
		if ctx.peekChar() == '0' {
			ctx.nextChar()
			isPunyCoded = true
		} else {
			hasWordSubSts = true
		}
	}
	var id strings.Builder
	for {
		for hasWordSubSts && unicode.IsLetter(ctx.peekChar()) {
			c = ctx.nextChar()
			var wordIdx int
			if unicode.IsLower(c) {
				wordIdx = int(c - 'a')
			} else {
				wordIdx = int(c - 'A')
				hasWordSubSts = false
			}
			if wordIdx >= ctx.numWords {
				return nil, fmt.Errorf("word index %d out of range (numWords=%d)", wordIdx, ctx.numWords)
			}
			id.WriteString(ctx.words[wordIdx])
		}
		if ctx.nextIf('0') {
			break
		}
		numChars := ctx.natural()
		if numChars < 0 {
			return nil, fmt.Errorf("invalid character count %d", numChars)
		}
		if isPunyCoded {
			ctx.nextIf('_')
		}
		if ctx.Pos+numChars > ctx.Size {
			return nil, fmt.Errorf("ran past end of buffer")
		}
		strRef := ctx.Data[ctx.Pos : ctx.Pos+numChars]
		if isPunyCoded {
			return nil, fmt.Errorf("punycode not implemented")
		} else {
			id.Write(strRef)
			wordStartPos := -1
			for idx := 0; idx <= len(strRef); idx++ {
				if idx < len(strRef) {
					c = rune(strRef[idx])
				} else {
					c = 0
				}
				if wordStartPos >= 0 && isWordEnd(c, rune(strRef[idx-1])) {
					if idx-wordStartPos >= 2 && ctx.numWords < maxWords {
						word := strRef[wordStartPos:idx]
						ctx.words[ctx.numWords] = string(word)
						ctx.numWords++
					}
					wordStartPos = -1
				}
				if wordStartPos < 0 && isWordStart(c) {
					wordStartPos = idx
				}
			}
		}
		ctx.Pos += numChars
		if !hasWordSubSts {
			break
		}
	}
	if id.Len() == 0 {
		return nil, fmt.Errorf("empty identifier")
	}
	ident := createNodeWithText(IdentifierKind, id.String())
	ctx.addSubstitution(ident)
	return ident, nil
}
