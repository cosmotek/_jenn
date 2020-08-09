package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/cosmotek/_jenn/hll/lexer"
)

func oneOf(a lexer.Token, b ...string) bool {
	for _, token := range b {
		if a.Type == token {
			return true
		}
	}

	return false
}

type E map[string]interface{}

type Walker struct {
	Tokens   []lexer.Token
	Position int

	Expressions []E
}

func (w *Walker) Next() lexer.Token {
	w.Position += 1
	if len(w.Tokens) < w.Position-1 {
		return lexer.Token{Type: lexer.EOF}
	}

	return w.Tokens[w.Position-1]
}

func (w *Walker) Walk() error {
	next := w.Next()

	switch next.Type {
	case lexer.APP:
		next = w.Next()
		if next.Type == lexer.IDENT {
			w.Expressions = append(w.Expressions, E{
				"appName": next.Literal,
			})
		}

		return w.Walk()

	case lexer.COMMENT:
		w.Expressions = append(w.Expressions, E{
			"comment": next.Literal,
		})

		return w.Walk()

	case lexer.ANNOTATION:
		next = w.Next()
		if next.Type == lexer.IDENT {
			w.Expressions = append(w.Expressions, E{
				"annotation": next.Literal,
			})
		}

		return w.Walk()

	// case lexer.ENUM:
	// 	name := ""
	// 	opts := []string{}

	// 	next = w.Next()
	// 	if next.Type == lexer.IDENT {
	// 		name = next.Literal
	// 		next = w.Next()
	// 		if next.Type != lexer.LBRACKET {
	// 			return errors.New("missing block")
	// 		}

	// 		if next.Type == lexer.RBRACKET {

	// 		}

	// 		// TODO WITH WHITESPACE SKIPPER
	// 	} else {
	// 		return errors.New("missing ident")
	// 	}

	// 	w.Expressions = append(w.Expressions, E{
	// 		"enum": E{
	// 			"name": name,
	// 			"opts": opts,
	// 		},
	// 	})

	// 	return w.Walk()

	case lexer.EOF:
		return nil

	default:
		return w.Walk()
	}
}

func main() {
	data, err := ioutil.ReadFile("hll/examples/shakenNotStirred.skm")
	if err != nil {
		panic(err)
	}

	lex := lexer.New(string(data))
	tokens := []lexer.Token{}

	for {
		tok, lit := lex.NextTokenWithLiteral()
		if tok != lexer.SPACE && tok != lexer.TAB {
			tokens = append(tokens, lexer.Token{
				Type:    tok,
				Literal: lit,
			})
		}

		if tok == lexer.EOF {
			break
		}
	}

	w := Walker{
		Tokens: tokens,
	}

	err = w.Walk()
	if err != nil {
		panic(err)
	} else {
		err = json.NewEncoder(os.Stdout).Encode(w.Expressions)
		if err != nil {
			panic(err)
		}
	}
}
