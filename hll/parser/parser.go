package parser

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cosmotek/_jenn/hll/lexer"
	"github.com/cosmotek/_jenn/ir"
)

type Parser struct {
	position uint64
	tokens   []lexer.Token

	bufferedComments    []string
	bufferedAnnotations []string

	model ir.ModelIR
}

func New(tokens ...lexer.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

/*
Parse walks the bank of tokens passed on construction. This parser
operates by trying to parse blocks (1-n lines in size) starting
with the expected keywords.

Parser tree structures can start with any of the following:
- app declaration
- enum declaration
- annotation
- comment
- type definition
*/
func (p *Parser) Parse() (ir.ModelIR, error) {
	for {
		tok := p.nextToken()
		if tok.Type == lexer.EOF {
			return p.model, nil
		}

		var err error
		switch tok.Type {
		case lexer.APP:
			err = p.ParseAppBlock()
			break

		case lexer.ENUM:
			err = p.ParseEnumBlock()
			break

		case lexer.ANNOTATION:
			err = p.ParseAnnotation()
			break

		case lexer.TYPE:
			err = p.ParseTypeBlock()
			break

		case lexer.COMMENT:
			p.bufferedComments = append(p.bufferedComments, tok.Literal)
			break

		// skip whitespace characters
		case lexer.NEWLINE, lexer.SPACE:
			break

		// handle any other tokens...
		default:
			return ir.ModelIR{}, fmt.Errorf("%d:%d: expected declaration, found '%s'", tok.Line, tok.ColumnStart, tok.Literal)
		}

		// check if err
		if err != nil {
			return ir.ModelIR{}, err
		}
	}
}

func (p *Parser) JSON() (string, error) {
	buff := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(buff).Encode(map[string]interface{}{
		"comments":    p.bufferedComments,
		"annotations": p.bufferedAnnotations,
		"position":    p.position,
	})
	if err != nil {
		return "", err
	}

	return string(buff.Bytes()), nil
}

func (p *Parser) flushBufferedComments() []string {
	oldComments := p.bufferedComments
	p.bufferedComments = []string{}

	return oldComments
}

func (p *Parser) nextToken() lexer.Token {
	if p.position > uint64(len(p.tokens)+1) {
		return lexer.Token{Type: lexer.EOF}
	}

	p.position++
	return p.tokens[p.position-1]
}

func (p *Parser) peekNextToken() lexer.Token {
	if p.position > uint64(len(p.tokens)+1) {
		return lexer.Token{Type: lexer.EOF}
	}

	return p.tokens[p.position]
}
