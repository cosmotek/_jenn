package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cosmotek/_jenn/hll/lexer"
	"github.com/cosmotek/_jenn/ir"
	"github.com/cosmotek/_jenn/ir/types"
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

	fields := []ir.Field{}

	// look for fields
	for {
		currTok := p.nextToken()
		typ := currTok.Type
		if typ == lexer.RBRACE {
			break
		}

		switch typ {
		case lexer.IDENT:
			fieldName := currTok.Literal
			sep := p.nextToken()
			if sep.Type != lexer.COLON {
				return fmt.Errorf("%d:%d: expected separator ':', found illegal '%s'", sep.Line, sep.ColumnStart, sep.Literal)
			}

			spaceOrType := p.nextToken()
			typeOfSpaceOrType := spaceOrType.Type
			if typeOfSpaceOrType != lexer.SPACE && typeOfSpaceOrType != lexer.IDENT {
				return fmt.Errorf("%d:%d: expected type name (ident), found illegal '%s'", spaceOrType.Line, spaceOrType.ColumnStart, spaceOrType.Literal)
			}

			typeName := spaceOrType.Literal
			if typeOfSpaceOrType == lexer.SPACE {
				typeTok := p.nextToken()
				typeOfSpaceOrType := typeTok.Type
				if typeOfSpaceOrType != lexer.SPACE && typeOfSpaceOrType != lexer.IDENT {
					return fmt.Errorf("%d:%d: expected type name (ident), found illegal '%s'", typeTok.Line, typeTok.ColumnStart, typeTok.Literal)
				}

				typeName = typeTok.Literal
			}

			for {
				next := p.nextToken()
				if next.Type == lexer.NEWLINE {
					fields = append(fields, ir.Field{
						Name:        fieldName,
						TypeOf:      types.CanonicalName(typeName), // todo resolve
						Description: strings.Join(p.flushBufferedComments(), "\n"),
					})

					break
				}

				if next.Type != lexer.SPACE && next.Type != lexer.TAB {
					return fmt.Errorf("%d:%d: expected newline, found illegal '%s'", next.Line, next.ColumnStart, next.Literal)
				}
			}
		case lexer.COMMENT:
			p.bufferedComments = append(p.bufferedComments, currTok.Literal)
			break
		case lexer.TAB, lexer.SPACE, lexer.NEWLINE:
			break
		default:
			return fmt.Errorf("%d:%d: expected field, found illegal '%s'", currTok.Line, currTok.ColumnStart, currTok.Literal)
		}
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
		Fields:      fields,
		//Fields
	})

	return nil
}
