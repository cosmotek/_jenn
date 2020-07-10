package parser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/cosmotek/_jenn/hll/lexer"
	"github.com/cosmotek/_jenn/ir"
)

func Parse(filename string) (ir.ModelIR, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return ir.ModelIR{}, err
	}

	model := ir.ModelIR{}
	lex := lexer.New(string(bytes))

	appName, description, err := parsePreamble(lex)
	if err != nil {
		return ir.ModelIR{}, err
	}

	model.Name = appName
	model.Description = description

	err = parseBody(lex, &model)
	if err != nil {
		return ir.ModelIR{}, err
	}

	return model, nil
}

func parsePreamble(lex *lexer.Lexer) (string, string, error) {
	comments := []string{}

	for {
		token, literal := lex.NextTokenWithLiteral()
		fmt.Println(token, literal)

		if token != lexer.APP && token != lexer.COMMENT && token != lexer.NEWLINE {
			return "", "", errors.New("missing preamble. schema must start with comment or preamble")
		}

		if token == lexer.COMMENT {
			comments = append(comments, literal)
		}

		if token == lexer.APP {
			spaceToken, _ := lex.NextTokenWithLiteral()
			nameToken, literal := lex.NextTokenWithLiteral()
			if spaceToken == lexer.SPACE && nameToken == lexer.IDENT {
				return literal, strings.Join(comments, ","), nil
			}

			return "", "", errors.New("syntax error, preamble incorrectly formed")
		}
	}
}

func parseBody(lex *lexer.Lexer, model *ir.ModelIR) error {
	for {
		token, literal := lex.NextTokenWithLiteral()
		fmt.Println(token, literal)

		if token != lexer.NEWLINE {
			switch token {
			case lexer.COMMENT:
				break
			case lexer.TYPE:
				newType := ir.Structure{}
				lex.NextTokenWithLiteral()
				identToken, literal := lex.NextTokenWithLiteral()
				spaceToken, _ := lex.NextTokenWithLiteral()
				braceToken, _ := lex.NextTokenWithLiteral()

				if identToken == lexer.IDENT && spaceToken == lexer.SPACE && braceToken == lexer.LBRACE {
					newType.Name = literal
				} else {
					return errors.New("invalid type declaration")
				}

				model.Types = append(model.Types, newType)
				break
			case lexer.ENUM:
				break
			case lexer.ANNOTATION:
				break
			}

			if token == lexer.EOF {
				return nil
			}
		}
	}
}
