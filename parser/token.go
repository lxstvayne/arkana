package parser

import "fmt"

type TokenType int

const (
	TOKENTYPE_NUMBER TokenType = iota
	TOKENTYPE_HEX_NUMBER

	TOKENTYPE_PLUS
	TOKENTYPE_MINUS
	TOKENTYPE_STAR
	TOKENTYPE_SLASH

	TOKENTYPE_LPAR // (
	TOKENTYPE_RPAR // )

	TOKENTYPE_EOF
)

func (tokenType TokenType) String() string {
	switch tokenType {
	case TOKENTYPE_NUMBER:
		return "NUMBER"
	case TOKENTYPE_HEX_NUMBER:
		return "HEX_NUMBER"
	case TOKENTYPE_PLUS:
		return "PLUS"
	case TOKENTYPE_MINUS:
		return "MINUS"
	case TOKENTYPE_STAR:
		return "STAR"
	case TOKENTYPE_SLASH:
		return "SLASH"
	case TOKENTYPE_EOF:
		return "EOF"
	default:
		return fmt.Sprintf("%d", tokenType)
	}
}

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
