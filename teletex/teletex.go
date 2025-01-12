// Decode Teletex (T.61) strings into native UTF-8 strings
package teletex

// Convert a single Teletex (T.61) character to its UTF-8 equivalent.
//
// Any code points that are unrecognized as Teletex will be returned unchanged
// to be interpreted directly as a UTF-8 code point. The rationale for this is
// that there is very high overlap between Teletex and ASCII and UTF-8. If
// there are any characters that I missed when writing this function, or if any
// new ones are added to Teletex, they are probably analogous.
func TeletexCharToUTF8(c byte) rune {
	switch c {
	case 0xA4:
		return '$'
	case 0xA6:
		return '#'
	case 0xA8:
		return '¤' // U+00A4
	case 0xB4:
		return '×' // U+00D7
	case 0xB8:
		return '÷' // U+00F7
	case 0xE0:
		return 'Ω' // U+2126
	case 0xE1:
		return 'Æ' // U+00C6
	case 0xE2:
		return 'Ð' // U+00D0
	case 0xE3:
		return 'ª' // U+00AA
	case 0xE4:
		return 'Ħ' // U+0126
	case 0xE6:
		return 'Ĳ' // U+0132
	case 0xE7:
		return 'Ŀ' // U+013F
	case 0xE8:
		return 'Ł' // U+0141
	case 0xE9:
		return 'Ø' // U+00D8
	case 0xEA:
		return 'Œ' // U+0152
	case 0xEB:
		return 'º' // U+00BA
	case 0xEC:
		return 'Þ' // U+00DE
	case 0xED:
		return 'Ŧ' // U+0166
	case 0xEE:
		return 'Ŋ' // U+014A
	case 0xEF:
		return 'ŉ' // U+0149
	case 0xF0:
		return 'ĸ' // U+0138
	case 0xF1:
		return 'æ' // U+00E6
	case 0xF2:
		return 'đ' // U+0111
	case 0xF3:
		return 'ð' // U+00F0
	case 0xF4:
		return 'ħ' // U+0127
	case 0xF5:
		return 'ı' // U+0131
	case 0xF6:
		return 'ĳ' // U+0133
	case 0xF7:
		return 'ŀ' // U+0140
	case 0xF8:
		return 'ł' // U+0142
	case 0xF9:
		return 'ø' // U+00F8
	case 0xFA:
		return 'œ' // U+0153
	case 0xFB:
		return 'ß' // U+00DF
	case 0xFC:
		return 'þ' // U+00FE
	case 0xFD:
		return 'ŧ' // U+0167
	case 0xFE:
		return 'ŋ' // U+014B

	// Diacritics
	case 0xC1:
		return '\u0300'
	case 0xC2:
		return '\u0301'
	case 0xC3:
		return '\u0302'
	case 0xC4:
		return '\u0303'
	case 0xC5:
		return '\u0304'
	case 0xC6:
		return '\u0306'
	case 0xC7:
		return '\u0307'
	case 0xC8:
		return '\u0308'
	case 0xC9:
		return '\u0308'
	case 0xCA:
		return '\u030A'
	case 0xCB:
		return '\u0327'
	case 0xCC:
		return '\u0332'
	case 0xCD:
		return '\u030B'
	case 0xCE:
		return '\u0328'
	case 0xCF:
		return '\u030C'
	default:
		return rune(c)
	}
}

// Determine whether a byte is a Teletex diacritic
func IsTeletexDiacritic(c byte) bool {
	return (c > 0xC0) && (c <= 0xCF)
}

type Teletex = []byte

// Convert a Teletex string (T.61 string) to a native Golang UTF-8 `string`
//
// A trailing diacritic is silently ignored.
func TeletexToUTF8(t Teletex) string {
	if len(t) <= 64 {
		is_ascii := true
		for _, v := range t {
			if v > 128 {
				is_ascii = false
				break
			}
		}
		if is_ascii {
			return string(t)
		}
	}
	runes := make([]rune, 0, len(t))
	var currentDiacritic rune = 0
	for _, v := range t {
		if currentDiacritic > 0 {
			// In UTF-8 the diacritic comes after the letter.
			// In Teletex, it comes before. We swap here.
			runes = append(runes, TeletexCharToUTF8(v))
			runes = append(runes, currentDiacritic)
			currentDiacritic = 0
			continue
		}
		if IsTeletexDiacritic(v) {
			currentDiacritic = TeletexCharToUTF8(v)
			continue
		}
		runes = append(runes, TeletexCharToUTF8(v))
	}
	return string(runes)
}

