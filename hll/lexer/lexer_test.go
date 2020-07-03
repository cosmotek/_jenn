package lexer

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/mgutz/ansi"
)

var red = ansi.ColorCode("red")
var reset = ansi.ColorCode("reset")

func printTable(expected []string, got []string) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Expected", "Got"})

	for i, expectedVal := range expected {
		gotVal := "N/A"
		if i < len(got) {
			gotVal = got[i]
		}

		if expectedVal != gotVal {
			expectedVal = red + expectedVal + reset
			gotVal = red + gotVal + reset
		}

		tw.AppendRow(table.Row{expectedVal, gotVal})
	}

	tw.SetStyle(table.StyleLight)

	// customize the style and change some stuff
	tw.Style().Format.Header = text.FormatLower
	tw.Style().Format.Row = text.FormatLower
	tw.Style().Format.Footer = text.FormatLower
	tw.Style().Options.SeparateColumns = false

	// render it
	fmt.Printf("Test diff:\n%s\n", tw.Render())
}

type tokenTest struct {
	Input          string
	ExpectedTokens []string
}

var nextTokenTests = map[string]tokenTest{
	"basic test": tokenTest{
		Input: `
app ShakenNotStirred

enum BeverageType {
	BEER,
	LIQUOR,
	WINE,
}

type Beverage {
	name: Name
	proof: Number
	type: BeverageType
}
`,
		ExpectedTokens: []string{
			NEWLINE,
			APP,
			IDENT,
			NEWLINE,
			NEWLINE,
			ENUM,
			IDENT,
			LBRACE,
			NEWLINE,
			IDENT,
			COMMA,
			NEWLINE,
			IDENT,
			COMMA,
			NEWLINE,
			IDENT,
			COMMA,
			NEWLINE,
			RBRACE,
			NEWLINE,
			NEWLINE,
			TYPE,
			IDENT,
			LBRACE,
			NEWLINE,
			IDENT,
			COLON,
			IDENT,
			NEWLINE,
			IDENT,
			COLON,
			IDENT,
			NEWLINE,
			IDENT,
			COLON,
			IDENT,
			NEWLINE,
			RBRACE,
		},
	},
}

func TestNextToken(t *testing.T) {
	for name, test := range nextTokenTests {
		lex := New(test.Input)

		t.Run(name, func(it *testing.T) {
			tokens := []string{}

			for {
				tok := lex.NextToken()
				tokens = append(tokens, tok)

				if tok == EOF {
					break
				}
			}

			if !cmp.Equal(tokens, test.ExpectedTokens) {
				printTable(test.ExpectedTokens, tokens)
				t.Fail()
				//t.Errorf("expected tokens '%v' but got '%v", test.ExpectedTokens, tokens)
			}
		})
	}
}
