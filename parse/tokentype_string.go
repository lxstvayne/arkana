// Code generated by "stringer --type TokenType --trimprefix TOKENTYPE_ ./parse"; DO NOT EDIT.

package parse

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TOKENTYPE_NUMBER-0]
	_ = x[TOKENTYPE_HEX_NUMBER-1]
	_ = x[TOKENTYPE_NONE-2]
	_ = x[TOKENTYPE_WORD-3]
	_ = x[TOKENTYPE_TEXT-4]
	_ = x[TOKENTYPE_IF-5]
	_ = x[TOKENTYPE_ELSE-6]
	_ = x[TOKENTYPE_WHILE-7]
	_ = x[TOKENTYPE_FOR-8]
	_ = x[TOKENTYPE_BREAK-9]
	_ = x[TOKENTYPE_CONTINUE-10]
	_ = x[TOKENTYPE_FUNC-11]
	_ = x[TOKENTYPE_RETURN-12]
	_ = x[TOKENTYPE_PLUS-13]
	_ = x[TOKENTYPE_MINUS-14]
	_ = x[TOKENTYPE_STAR-15]
	_ = x[TOKENTYPE_SLASH-16]
	_ = x[TOKENTYPE_EQ-17]
	_ = x[TOKENTYPE_EQEQ-18]
	_ = x[TOKENTYPE_EXCL-19]
	_ = x[TOKENTYPE_EXCLEQ-20]
	_ = x[TOKENTYPE_LT-21]
	_ = x[TOKENTYPE_LTEQ-22]
	_ = x[TOKENTYPE_GT-23]
	_ = x[TOKENTYPE_GTEQ-24]
	_ = x[TOKENTYPE_BAR-25]
	_ = x[TOKENTYPE_BARBAR-26]
	_ = x[TOKENTYPE_AMP-27]
	_ = x[TOKENTYPE_AMPAMP-28]
	_ = x[TOKENTYPE_LPAR-29]
	_ = x[TOKENTYPE_RPAR-30]
	_ = x[TOKENTYPE_LBRACE-31]
	_ = x[TOKENTYPE_RBRACE-32]
	_ = x[TOKENTYPE_LBRACKET-33]
	_ = x[TOKENTYPE_RBRACKET-34]
	_ = x[TOKENTYPE_COMMA-35]
	_ = x[TOKENTYPE_SEMICOLON-36]
	_ = x[TOKENTYPE_EOF-37]
}

const _TokenType_name = "NUMBERHEX_NUMBERNONEWORDTEXTIFELSEWHILEFORBREAKCONTINUEFUNCRETURNPLUSMINUSSTARSLASHEQEQEQEXCLEXCLEQLTLTEQGTGTEQBARBARBARAMPAMPAMPLPARRPARLBRACERBRACELBRACKETRBRACKETCOMMASEMICOLONEOF"

var _TokenType_index = [...]uint8{0, 6, 16, 20, 24, 28, 30, 34, 39, 42, 47, 55, 59, 65, 69, 74, 78, 83, 85, 89, 93, 99, 101, 105, 107, 111, 114, 120, 123, 129, 133, 137, 143, 149, 157, 165, 170, 179, 182}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
