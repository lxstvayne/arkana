package parser

import (
	"strings"
	"unicode"
)

const (
	OPERATOR_CHARS = "+-*/()=<>!&|"
)

var (
	OPERATORS = map[string]TokenType{
		"+": TOKENTYPE_PLUS,
		"-": TOKENTYPE_MINUS,
		"*": TOKENTYPE_STAR,
		"/": TOKENTYPE_SLASH,
		"(": TOKENTYPE_LPAR,
		")": TOKENTYPE_RPAR,
		"=": TOKENTYPE_EQ,
		"<": TOKENTYPE_LT,
		">": TOKENTYPE_GT,

		"!": TOKENTYPE_EXCL,
		"&": TOKENTYPE_AMP,
		"|": TOKENTYPE_BAR,

		"==": TOKENTYPE_EQEQ,
		"!=": TOKENTYPE_EXCLEQ,
		"<=": TOKENTYPE_LTEQ,
		">=": TOKENTYPE_GTEQ,

		"&&": TOKENTYPE_AMPAMP,
		"||": TOKENTYPE_BARBAR,
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
		current, _ := lexer.peek(0)

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
		} else if current == rune('"') {
			lexer.tokenizeText()
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
	current, _ := lexer.peek(0)
	for {
		if current == rune('.') {
			if strings.Contains(buf.String(), ".") {
				panic("Invalid float number")
			}
		} else if !unicode.IsDigit(current) {
			break
		}

		buf.WriteRune(current)
		current, _ = lexer.next()
	}
	lexer.addToken(TOKENTYPE_NUMBER, []rune(buf.String()))
}

func (lexer *Lexer) tokenizeWord() {
	var buf strings.Builder
	current, _ := lexer.peek(0)
	for {
		if !(unicode.IsLetter(current) || unicode.IsDigit(current) || current == rune('_')) {
			break
		}

		buf.WriteRune(current)
		current, _ = lexer.next()
	}
	switch buf.String() {
	case "print":
		lexer.addToken(TOKENTYPE_PRINT, nil)
	case "if":
		lexer.addToken(TOKENTYPE_IF, nil)
	case "else":
		lexer.addToken(TOKENTYPE_ELSE, nil)
	default:
		lexer.addToken(TOKENTYPE_WORD, []rune(buf.String()))
	}
}

func (lexer *Lexer) tokenizeText() {
	lexer.next() // skip "
	var buf strings.Builder
	current, _ := lexer.peek(0)
	for {
		if current == rune('\\') {
			current, _ = lexer.next()
			switch current {
			case rune('"'):
				current, _ = lexer.next()
				buf.WriteRune(rune('"'))
				continue
			case rune('n'):
				current, _ = lexer.next()
				buf.WriteRune(rune('\n'))
				continue
			case rune('t'):
				current, _ = lexer.next()
				buf.WriteRune(rune('\t'))
				continue
			}
			buf.WriteRune(rune('\\'))
			continue
		}
		if current == rune('"') {
			break
		}

		buf.WriteRune(current)
		current, _ = lexer.next()
	}
	lexer.next() // skip closing "
	lexer.addToken(TOKENTYPE_TEXT, []rune(buf.String()))

}

func (lexer *Lexer) tokenizeHexNumber() {
	var buf strings.Builder
	current, _ := lexer.peek(0)
	for {
		if !unicode.IsDigit(current) && !strings.ContainsRune("abcdef", unicode.ToLower(current)) {
			break
		}

		buf.WriteRune(current)
		current, _ = lexer.next()
	}
	lexer.addToken(TOKENTYPE_HEX_NUMBER, []rune(buf.String()))
}

func (lexer *Lexer) tokenizeOperator() {
	current, _ := lexer.peek(0)
	if current == '/' {
		next, _ := lexer.peek(1)
		if next == '/' {
			lexer.next()
			lexer.next()
			lexer.tokenizeComment()
			return
		} else if next == '*' {
			lexer.next()
			lexer.next()
			lexer.tokenizeMultilineComment()
			return
		}
	}

	var buf strings.Builder
	for {
		text := buf.String()
		if _, ok := OPERATORS[text+string(current)]; !ok && text != "" {
			lexer.addToken(OPERATORS[text], nil)
			return
		}
		buf.WriteRune(current)
		current, _ = lexer.next()
	}
}

func (lexer *Lexer) tokenizeComment() {
	current, isEOF := lexer.peek(0)

	for {
		if isEOF || strings.ContainsRune("\r\n", current) {
			return
		}

		current, isEOF = lexer.next()
	}
}

func (lexer *Lexer) tokenizeMultilineComment() {
	current, isEOF := lexer.peek(0)
	for {
		if isEOF {
			panic("Missing close tag")
		}
		next, _ := lexer.peek(1)
		if current == '*' && next == '/' {
			break
		}

		current, isEOF = lexer.next()
	}

	lexer.next() // *
	lexer.next() // /
}

func (lexer *Lexer) peek(relativePos int) (rune, bool) {
	pos := lexer.position + relativePos
	if pos >= lexer.lenght {
		return 0, true
	}

	return lexer.input[pos], false
}

func (lexer *Lexer) next() (rune, bool) {
	lexer.position += 1
	return lexer.peek(0)
}
