package lexer_test

import (
	"testing"

	"github.com/Pshimaf-Git/corvus/internal/lexer"
	"github.com/Pshimaf-Git/corvus/internal/token"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		const code = "fn main() {\n return 1 \n}"
		l := lexer.New(code)
		assert.NotZero(t, l)
		assert.Equal(t, code, l.Source())
	})
}

func TestLexer(t *testing.T) {
	t.Run("NextToken", func(t *testing.T) {
		testCases := []struct {
			name string
			code string
			want []token.Token
		}{
			{
				name: "simple function",
				code: `
				fn main() {
				  s = "string data";
					return sum(1,2);
				}

				fn sum(a int, b int) {
					res = a + b;
					return res;
				}`,
				want: []token.Token{
					token.WithLiteral(token.FN, "fn"),
					token.WithLiteral(token.IDENTIFIER, "main"),
					token.WithLiteral(token.LPAREN, "("),
					token.WithLiteral(token.RPAREN, ")"),
					token.WithLiteral(token.LBRACE, "{"),
					token.WithLiteral(token.IDENTIFIER, "s"),
					token.WithLiteral(token.ASIGN, "="),
					token.WithLiteral(token.STRING, "string data"),
					token.WithLiteral(token.SEMICOLON, ";"),

					token.WithLiteral(token.RETUNR, "return"),
					token.WithLiteral(token.IDENTIFIER, "sum"),
					token.WithLiteral(token.LPAREN, "("),
					token.WithLiteral(token.INTEGER, "1"),
					token.WithLiteral(token.COMMA, ","),
					token.WithLiteral(token.INTEGER, "2"),
					token.WithLiteral(token.RPAREN, ")"),
					token.WithLiteral(token.SEMICOLON, ";"),
					token.WithLiteral(token.RBRACE, "}"),

					token.WithLiteral(token.FN, "fn"),
					token.WithLiteral(token.IDENTIFIER, "sum"),
					token.WithLiteral(token.LPAREN, "("),
					token.WithLiteral(token.IDENTIFIER, "a"),
					token.WithLiteral(token.INT, "int"),
					token.WithLiteral(token.COMMA, ","),
					token.WithLiteral(token.IDENTIFIER, "b"),
					token.WithLiteral(token.INT, "int"),
					token.WithLiteral(token.RPAREN, ")"),
					token.WithLiteral(token.LBRACE, "{"),
					token.WithLiteral(token.IDENTIFIER, "res"),
					token.WithLiteral(token.ASIGN, "="),
					token.WithLiteral(token.IDENTIFIER, "a"),
					token.WithLiteral(token.PLUS, "+"),
					token.WithLiteral(token.IDENTIFIER, "b"),
					token.WithLiteral(token.SEMICOLON, ";"),
					token.WithLiteral(token.RETUNR, "return"),
					token.WithLiteral(token.IDENTIFIER, "res"),
					token.WithLiteral(token.SEMICOLON, ";"),
					token.WithLiteral(token.RBRACE, "}"),

					token.New(token.EOF),
				},
			},

			{
				name: "bad string declare",
				code: `
				s = "bad
				`,
				want: []token.Token{
					token.WithLiteral(token.IDENTIFIER, "s"),
					token.WithLiteral(token.ASIGN, "="),
					token.New(token.BAD_SYNTAX),

					token.New(token.EOF),
				},
			},
		}

		for _, tt := range testCases {
			t.Run(tt.name, func(t *testing.T) {
				l := lexer.New(tt.code)
				for _, wantToken := range tt.want {
					gotToken := l.NextToken()
					assert.Equal(t, wantToken, gotToken)
				}
			})
		}
	})

	t.Run("Source", func(t *testing.T) {
		const code = "fn main() {\n\n\n\r\t return 0\n\n\t}\n\n\n"
		l := lexer.New(code)

		assert.Equal(t, code, l.Source())
	})
}
