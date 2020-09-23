package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cosmotek/_jenn/hll/lexer"
	"github.com/cosmotek/_jenn/ir"
)

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
