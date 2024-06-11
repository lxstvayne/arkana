package parse

type TokenType int

const (
	TOKENTYPE_NUMBER TokenType = iota
	TOKENTYPE_HEX_NUMBER
	TOKENTYPE_NONE
	TOKENTYPE_WORD
	TOKENTYPE_TEXT

	// keyword
	TOKENTYPE_IF
	TOKENTYPE_ELSE
	TOKENTYPE_WHILE
	TOKENTYPE_FOR
	TOKENTYPE_BREAK
	TOKENTYPE_CONTINUE
	TOKENTYPE_FUNC
	TOKENTYPE_RETURN

	TOKENTYPE_PLUS
	TOKENTYPE_MINUS
	TOKENTYPE_STAR
	TOKENTYPE_SLASH
	TOKENTYPE_EQ
	TOKENTYPE_EQEQ
	TOKENTYPE_EXCL
	TOKENTYPE_EXCLEQ
	TOKENTYPE_LT
	TOKENTYPE_LTEQ
	TOKENTYPE_GT
	TOKENTYPE_GTEQ

	TOKENTYPE_BAR
	TOKENTYPE_BARBAR
	TOKENTYPE_AMP
	TOKENTYPE_AMPAMP

	TOKENTYPE_LPAR      // (
	TOKENTYPE_RPAR      // )
	TOKENTYPE_LBRACE    // {
	TOKENTYPE_RBRACE    // }
	TOKENTYPE_LBRACKET  // [
	TOKENTYPE_RBRACKET  // ]
	TOKENTYPE_COMMA     // ,
	TOKENTYPE_SEMICOLON // ;

	TOKENTYPE_EOF
)

type Token struct {
	tokenType TokenType
	text      []rune
}

func NewToken(tokenType TokenType, text []rune) *Token {
	return &Token{
		tokenType: tokenType,
		text:      text,
	}
}

func (token *Token) TokenType() TokenType {
	return token.tokenType
}

func (token *Token) Text() []rune {
	return token.text
}

var (
	TOKEN_EOF = NewToken(TOKENTYPE_EOF, nil)
)
