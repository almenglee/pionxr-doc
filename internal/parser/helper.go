package parser

func isLetter(ch rune) bool  { return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_' }
func lower(ch rune) rune     { return ('a' - 'A') | ch }
func isDecimal(ch rune) bool { return '0' <= ch && ch <= '9' }
func isIdent(str string) bool {
	if str == "" {
		return false
	}
	if !isLetter(rune(str[0])) {
		return false
	}

	for _, ch := range str {
		if !isLetter(rune(ch)) && !isDecimal(rune(ch)) {
			return false
		}
	}

	return true
}
