package parser

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/cosmotek/_jenn/hll/lexer"
	"github.com/cosmotek/_jenn/ir"
)

type ParserOutput struct{}

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

func (p *Parser) ParseAppBlock() error {
	if p.model.Name != "" {
		return errors.New("duplicate app preamble")
	}

	nextTok := p.nextToken()
	lastTok := p.nextToken()

	if nextTok.Type == lexer.SPACE && lastTok.Type == lexer.IDENT {
		p.model.Description = strings.Join(p.flushBufferedComments(), "\n")
		p.model.Name = lastTok.Literal

		return nil
	}

	return fmt.Errorf("invalid app preamble '%s%s'", nextTok.Literal, lastTok.Literal)
}

func (p *Parser) ParseEnumBlock() error {
	startingSeq := []string{lexer.SPACE, lexer.IDENT, lexer.SPACE, lexer.LBRACE}
	enumName := ""
	vals := map[string]string{}

	for i, tok := range startingSeq {
		currTok := p.nextToken()

		if tok != currTok.Type {
			return fmt.Errorf("invalid enum block, unexpected token %s->'%s' at position %v", currTok.Type, currTok.Literal, i)
		}

		if i == 1 {
			enumName = currTok.Literal
		}
	}

	lineStarted := false
	bufferedComments := []string{}

	for {
		currTok := p.nextToken()
		if currTok.Type == lexer.RBRACE {
			break
		}

		switch currTok.Type {
		case lexer.SPACE, lexer.TAB:
			if lineStarted {
				return errors.New("whitespace not allowed at end of line")
			}

		case lexer.IDENT:
			lineStarted = true
			vals[currTok.Literal] = strings.Join(bufferedComments, "\n")

			break

		// case lexer.COMMA:
		// 	if !lineStarted || lineEnded {
		// 		return fmt.Errorf("invalid enum block, unexpected token %s->'%s'", currTok.Type, currTok.Literal)
		// 	}

		case lexer.NEWLINE:
			lineStarted = false
			break

		case lexer.COMMENT:
			bufferedComments = append(bufferedComments, currTok.Literal)
			break

		default:
			return fmt.Errorf("invalid enum block, unexpected token %s->'%s'", currTok.Type, currTok.Literal)
		}
	}

	p.model.Enums = append(p.model.Enums, ir.Enum{
		Description: strings.Join(p.flushBufferedComments(), "\n"),
		Name:        enumName,
		Values:      vals,
	})

	return nil
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
