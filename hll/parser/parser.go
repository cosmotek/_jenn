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

		switch tok.Type {
		case lexer.APP:
			err := p.ParseAppBlock()
			if err != nil {
				return ir.ModelIR{}, err
			}

			break

		case lexer.ENUM:
			err := p.ParseEnumBlock()
			if err != nil {
				return ir.ModelIR{}, err
			}

			break

		case lexer.ANNOTATION:
			break
		case lexer.TYPE:
			break
		case lexer.COMMENT:
			p.bufferedComments = append(p.bufferedComments, tok.Literal)
			break

		// skip whitespace characters
		case lexer.NEWLINE, lexer.SPACE:
			break

		// handle any other tokens...
		default:
			return ir.ModelIR{}, fmt.Errorf("block starts with unacceptable token: %s->'%s'", tok.Type, tok.Literal)
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
