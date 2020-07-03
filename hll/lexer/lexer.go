package lexer

import "log"

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers & Literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Operators
	ASSIGN = "ASSIGN"

	// Delimiters
	COMMA   = "COMMA"
	COLON   = "COLON"
	NEWLINE = "NEWLINE"

	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	LBRACE = "LBRACE"
	RBRACE = "RBRACE"

	// Keywords
	TYPE = "TYPE"
	ENUM = "ENUM"
	APP  = "APP"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		log.Println("EOF")
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() string {
	defer l.readChar()
	token := func() string {
		switch l.ch {
		case '\n':
			return NEWLINE
		case '=':
			return ASSIGN
		case ',':
			return COMMA
		case ':':
			return COLON
		case '(':
			return LPAREN
		case ')':
			return RPAREN
		case 0:
			return EOF
		default:
			if isLetter(l.ch) {
				return l.readIdentifier()
			} else {
				return ILLEGAL
			}
		}
	}()

	return token
}

func (l *Lexer) readIdentifier() string {
	pos := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	log.Println(l.input[pos:l.position])

	return IDENT
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
