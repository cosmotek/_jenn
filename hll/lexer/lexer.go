package lexer

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	SPACE   = "SPACE"
	TAB     = "TAB"

	// Identifiers & Literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"
	TRUE   = "TRUE"
	FALSE  = "FALSE"

	// Operators
	ASSIGN     = "ASSIGN"
	NULLABLE   = "NULLABLE"
	ANNOTATION = "ANNOTATION"
	COMMENT    = "COMMENT"

	// Delimiters
	COMMA   = "COMMA"
	PERIOD  = "PERIOD"
	COLON   = "COLON"
	NEWLINE = "NEWLINE"

	LPAREN   = "LPAREN"
	RPAREN   = "RPAREN"
	LBRACE   = "LBRACE"
	RBRACE   = "RBRACE"
	LBRACKET = "LBRACKET"
	RBRACKET = "RBRACKET"

	// Keywords
	TYPE      = "TYPE"
	ENUM      = "ENUM"
	APP       = "APP"
	NAMESPACE = "NAMESPACE"
	SELECTOR  = "SELECTOR"

	// Selector Keywords
	FULLTEXT    = "FULLTEXT"
	EXACT       = "EXACT"
	EXACT_LARGE = "EXACT_LARGE"
	RANGE       = "RANGE"
	TERM        = "TERM"
	TRIGRAM     = "TRIGRAM"
)

var keywords = map[string]string{
	"type": TYPE,
	"enum": ENUM,
	"app":  APP,

	"true":  TRUE,
	"false": FALSE,

	"namespace": NAMESPACE,
	"selector":  SELECTOR,
	"fulltext":  FULLTEXT,
	"exact":     EXACT,
	"range":     RANGE,
}

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
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextTokenWithLiteral() (string, string) {
	char := l.ch
	token := func() string {
		switch l.ch {
		case ' ':
			return SPACE
		case '\n':
			return NEWLINE
		case '\t':
			return TAB
		case '=':
			return ASSIGN
		case ',':
			return COMMA
		case '?':
			return NULLABLE
		case '@':
			return ANNOTATION
		case ':':
			return COLON
		case '(':
			return LPAREN
		case ')':
			return RPAREN
		case '{':
			return LBRACE
		case '}':
			return RBRACE
		case '[':
			return LBRACKET
		case ']':
			return RBRACKET
		case 0:
			return EOF
		default:
			return ILLEGAL
		}
	}()

	if l.isComment(l.ch) {
		return COMMENT, l.readComment()
	}

	if isLetter(l.ch) {
		return l.readIdentifier()
	}

	l.readChar()
	return token, string(char)
}

func (l *Lexer) readIdentifier() (string, string) {
	pos := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return lookupIdent(l.input[pos:l.position]), l.input[pos:l.position]
}

func lookupIdent(identifier string) string {
	if token, ok := keywords[identifier]; ok {
		return token
	}

	return IDENT
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) isComment(ch byte) bool {
	return ch == '/' && l.input[l.position+1] == '/' || ch == '#'
}

func (l *Lexer) readComment() string {
	pos := l.position

	for {
		l.readChar()
		if l.ch == '\n' {
			return l.input[pos:l.position]
		}
	}
}

type Token struct {
	Type    string
	Literal string
}

func (l *Lexer) Tokens() []Token {
	tokens := []Token{}
	for {
		tok, lit := l.NextTokenWithLiteral()
		if tok == EOF {
			return tokens
		}

		tokens = append(tokens, Token{
			Type:    tok,
			Literal: lit,
		})
	}
}
