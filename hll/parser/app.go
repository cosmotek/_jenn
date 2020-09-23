package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cosmotek/_jenn/hll/lexer"
)

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
