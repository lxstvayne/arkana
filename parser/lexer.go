package parser

import (
	"strings"
	"unicode"
)

const (
	OPERATOR_CHARS = "+-*/()="
)

var (
	OPERATOR_TOKENS = [...]TokenType{
		TOKENTYPE_PLUS, TOKENTYPE_MINUS, TOKENTYPE_STAR, TOKENTYPE_SLASH,
		TOKENTYPE_LPAR, TOKENTYPE_RPAR, TOKENTYPE_EQ,
	}
)

type Lexer struct {
	input    []rune
	tokens   []*Token
	position int
	lenght   int
}

func NewLexer(input string) *Lexer {
	runeInput := []rune(input)
	lexer := &Lexer{
		input:  runeInput,
		lenght: len(runeInput),
	}
	return lexer
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

func (lexer *Lexer) tokenizeNumber() {
	var buf strings.Builder
	current := lexer.peek(0)
	for {
		if current == rune('.') {
			if strings.Contains(buf.String(), ".") {
				panic("Invalid float number")
			}
		} else if !unicode.IsDigit(current) {
			break
		}

		buf.WriteRune(current)
		current = lexer.next()
	}
	lexer.addToken(TOKENTYPE_NUMBER, []rune(buf.String()))
}

func (lexer *Lexer) tokenizeWord() {
	var buf strings.Builder
	current := lexer.peek(0)
	for {
		if !(unicode.IsLetter(current) || unicode.IsDigit(current) || current == rune('_')) {
			break
		}

		buf.WriteRune(current)
		current = lexer.next()
	}
	lexer.addToken(TOKENTYPE_WORD, []rune(buf.String()))
}

func (lexer *Lexer) tokenizeHexNumber() {
	var buf strings.Builder
	current := lexer.peek(0)
	for {
		if !unicode.IsDigit(current) && !strings.ContainsRune("abcdef", unicode.ToLower(current)) {
			break
		}

		buf.WriteRune(current)
		current = lexer.next()
	}
	lexer.addToken(TOKENTYPE_HEX_NUMBER, []rune(buf.String()))
}

func (lexer *Lexer) tokenizeOperator() {
	current := lexer.peek(0)
	position := strings.IndexRune(OPERATOR_CHARS, current)
	lexer.addToken(OPERATOR_TOKENS[position], nil)
	lexer.next()
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
