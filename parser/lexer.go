package parser

import (
	"strings"
	"unicode"

	"golang.org/x/exp/slices"
)

const (
	OPERATOR_CHARS = "+-*/()"
)

var (
	OPERATOR_TOKENS = [...]TokenType{
		TOKENTYPE_PLUS, TOKENTYPE_MINUS, TOKENTYPE_STAR, TOKENTYPE_SLASH,
		TOKENTYPE_LPAR, TOKENTYPE_RPAR,
	}
)

type Lexer struct {
	input    []rune
	tokens   []*Token
	lenght   int
	position int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  []rune(input),
		lenght: len(input),
		tokens: []*Token{},
	}
}

func (lexer *Lexer) Tokenize() []*Token {
	for {
		current := lexer.peek(0)

		if lexer.position >= lexer.lenght {
			break
		}

		if unicode.IsDigit(current) {
			lexer.tokenizeNumber()
		} else if unicode.IsLetter(current) {
			lexer.tokenizeWord()
		} else if current == rune('#') {
			lexer.next()
			lexer.tokenizeHexNumber()
		} else if strings.ContainsRune(OPERATOR_CHARS, current) {
			lexer.tokenizeOperator()
		} else {
			lexer.next()
		}
	}
	return lexer.tokens
}

func (lexer *Lexer) addToken(tokenType TokenType, text []rune) {
	lexer.tokens = append(lexer.tokens, NewToken(tokenType, text))
}

func (lexer *Lexer) peek(relativePos int) rune {
	pos := lexer.position + relativePos
	if pos >= lexer.lenght {
		return rune(0)
	}

	return lexer.input[pos]
}

func (lexer *Lexer) next() rune {
	lexer.position += 1
	return lexer.peek(0)
}

func (lexer *Lexer) tokenizeNumber() {
	var buf []rune
	current := lexer.peek(0)
	for {
		if current == rune('.') {
			if slices.Contains(buf, rune('.')) {
				panic("Invalid float number")
			}
		} else if !unicode.IsDigit(current) {
			break
		}

		buf = append(buf, current)
		current = lexer.next()
	}
	lexer.addToken(TOKENTYPE_NUMBER, buf)
}

func (lexer *Lexer) tokenizeWord() {
	var buf []rune
	current := lexer.peek(0)
	for {
		if !(unicode.IsLetter(current) || unicode.IsDigit(current) || current == rune('_')) {
			break
		}

		buf = append(buf, current)
		current = lexer.next()
	}
	lexer.addToken(TOKENTYPE_WORD, buf)
}

func (lexer *Lexer) tokenizeHexNumber() {
	var buf []rune
	current := lexer.peek(0)
	for {
		if !unicode.IsDigit(current) && !strings.ContainsRune("abcdef", unicode.ToLower(current)) {
			break
		}

		buf = append(buf, current)
		current = lexer.next()
	}
	lexer.addToken(TOKENTYPE_HEX_NUMBER, buf)
}

func (lexer *Lexer) tokenizeOperator() {
	position := strings.IndexRune(OPERATOR_CHARS, lexer.peek(0))
	lexer.addToken(OPERATOR_TOKENS[position], nil)
	lexer.next()
}
