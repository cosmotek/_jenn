package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cosmotek/_jenn/hll/lexer"
	"github.com/cosmotek/_jenn/ir"
)

func contains(list []string, item string) bool {
	for _, elem := range list {
		if elem == item {
			return true
		}
	}

	return false
}

func (p *Parser) ParseTypeBlock() error {
	startingDef := []string{lexer.SPACE, lexer.IDENT}
	typeName := ""

	// get name
	for i, tok := range startingDef {
		currTok := p.nextToken()

		if tok != currTok.Type {
			return fmt.Errorf("invalid type block head, unexpected token %s->'%s' at position %v", currTok.Type, currTok.Literal, i)
		}

		if i == 1 {
			typeName = currTok.Literal
		}
	}

	formFields := []string{}

	// check for form in def
	if p.peekNextToken().Type == lexer.LPAREN {
		// move position +1
		p.nextToken()

		breakAdded := false
		elemAdded := false

		// scan for form fields until rparen
		for {
			currTok := p.nextToken()
			if currTok.Type == lexer.RPAREN {
				break
			}

			// TODO add checks for optional operator and default values

			if !elemAdded {
				if currTok.Type != lexer.IDENT {
					return errors.New("expected ident")
				}

				if contains(formFields, currTok.Literal) {
					return fmt.Errorf("duplicate type form field '%s'", currTok.Literal)
				}

				formFields = append(formFields, currTok.Literal)
				elemAdded = true
			} else if elemAdded && !breakAdded {
				if currTok.Type != lexer.COMMA {
					return fmt.Errorf("expected comma, found literal %s", currTok.Type)
				}

				breakAdded = true
				nextTok := p.peekNextToken()

				// handle optional space and move position if found
				if nextTok.Type == lexer.SPACE {
					p.nextToken()
				}

				// reset state
				elemAdded = false
				breakAdded = false
			}
		}
	}

	// confirm opening of block using seq
	middleDef := []string{lexer.SPACE, lexer.LBRACE}

	for i, tok := range middleDef {
		currTok := p.nextToken()

		if tok != currTok.Type {
			return fmt.Errorf("invalid type block middle, unexpected token %s->'%s' at position %v", currTok.Type, currTok.Literal, i)
		}
	}

	// fields := map[string]string{}

	// look for fields
	for {
		currTok := p.nextToken()
		if currTok.Type == lexer.RBRACE {
			break
		}

		// TODO: parse fields
	}

	// add form fields to model
	formFieldIR := []ir.Field{}
	for _, fieldName := range formFields {
		formFieldIR = append(formFieldIR, ir.Field{
			Name: fieldName,
			// TODO future pull field options and typeOf from field in parent type
		})
	}

	p.model.Forms = append(p.model.Forms, ir.Structure{
		Name:   typeName,
		Fields: formFieldIR,
	})

	// add type to model
	p.model.Types = append(p.model.Types, ir.Structure{
		Description: strings.Join(p.flushBufferedComments(), "\n"),
		Name:        typeName + strings.Join(formFields, ","),
		//Fields
	})

	return nil
}
