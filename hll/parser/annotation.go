package parser

import (
	"fmt"

	"github.com/cosmotek/_jenn/hll/lexer"
)

func (p *Parser) ParseAnnotation() error {
	nextTok := p.nextToken()

	if nextTok.Type == lexer.IDENT {
		p.bufferedAnnotations = append(p.bufferedAnnotations, nextTok.Literal)

		for {
			currTok := p.nextToken()
			if currTok.Type == lexer.NEWLINE || currTok.Type == lexer.EOF {
				return nil
			}
		}
	}

	return fmt.Errorf("invalid annotation '%s'", nextTok.Literal)
}
