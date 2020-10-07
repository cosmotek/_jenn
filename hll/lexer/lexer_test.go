package lexer

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/mgutz/ansi"
)

var red = ansi.ColorCode("red")
var reset = ansi.ColorCode("reset")

func printTable(expected []string, got []string, literals []string) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Expected", "Got", "Literal"})

	for i, expectedVal := range expected {
		gotVal := ""
		if i < len(got) {
			gotVal = got[i]
		}

		litVal := ""
		if i < len(literals) {
			litVal = literals[i]
		}

		litVal = strings.Replace(litVal, "\n", "\\n", -1)
		litVal = strings.Replace(litVal, "\t", "\\t", -1)
		litVal = strings.Replace(litVal, " ", "\\s", -1)

		if expectedVal != gotVal {
			expectedVal = red + expectedVal + reset
			gotVal = red + gotVal + reset
		}

		tw.AppendRow(table.Row{expectedVal, gotVal, litVal})
	}

	if len(got) > len(expected) {
		diff := len(got) - len(expected)
		elems := got[len(got)-diff:]

		for i, gotVal := range elems {
			litVal := ""
			if i < len(literals) {
				litVal = literals[i]
			}

			litVal = strings.Replace(litVal, "\n", "\\n", -1)
			litVal = strings.Replace(litVal, "\t", "\\t", -1)
			litVal = strings.Replace(litVal, " ", "\\s", -1)

			gotVal = red + gotVal + reset
			tw.AppendRow(table.Row{"", gotVal, litVal})
		}
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
	"basic test": {
		Input: `
app ShakenNotStirred

// enums are cool
enum BeverageType {
	BEER,
	LIQUOR,
	WINE,
}

type Beverage {
	name: ?Name
	proof: Number
	typeOf: BeverageType
}
`,
		ExpectedTokens: []string{
			NEWLINE,
			APP,
			IDENT,
			NEWLINE,
			NEWLINE,
			COMMENT,
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
			NULLABLE,
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
			NEWLINE,
			EOF,
		},
	},
}

func TestNextToken(t *testing.T) {
	for name, test := range nextTokenTests {
		lex := New(test.Input)

		t.Run(name, func(it *testing.T) {
			tokens := []string{}
			literals := []string{}

			for {
				tok, lit := lex.NextTokenWithLiteral()
				if tok != SPACE && tok != TAB {
					tokens = append(tokens, tok)
					literals = append(literals, lit)
				}

				if tok == EOF {
					break
				}
			}

			if !cmp.Equal(tokens, test.ExpectedTokens) {
				printTable(test.ExpectedTokens, tokens, literals)
				t.Error(cmp.Diff(tokens, test.ExpectedTokens))
			}
		})
	}
}

type lineCountTest struct {
	Input         string
	ExpectedLines int
}

var lineCountTests = map[string]lineCountTest{
	"basic test": {
		Input: `app ShakenNotStirred

// enums are cool
enum BeverageType {
	BEER,
	LIQUOR,
	WINE,
}

type Beverage {
	name: ?Name
	proof: Number
	typeOf: BeverageType
}
`,
		ExpectedLines: 14,
	},
}

func TestLineCounts(t *testing.T) {
	for name, test := range lineCountTests {
		lex := New(test.Input)

		t.Run(name, func(it *testing.T) {
			// for {
			// 	tok, _ := lex.NextTokenWithLiteral()
			// 	if tok == EOF {
			// 		break
			// 	}
			// }

			t.Logf("%v", lex.Tokens())

			if !cmp.Equal(lex.LineCount(), test.ExpectedLines) {
				t.Error(cmp.Diff(lex.LineCount(), test.ExpectedLines))
			}
		})
	}
}
