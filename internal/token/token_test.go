package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenType_String(t *testing.T) {
	testCases := []struct {
		name      string
		tokenType TokenType
		want      string
	}{
		{
			name:      "left paren '('",
			tokenType: LPAREN,
			want:      "(",
		},
		{
			name:      "right paren ')'",
			tokenType: RPAREN,
			want:      ")",
		},
		{
			name:      "and '('",
			tokenType: AND,
			want:      "and",
		},
		{
			name:      "bool",
			tokenType: BOOL,
			want:      "<bool>",
		},
		{
			name:      "integer",
			tokenType: INTEGER,
			want:      "<integer>",
		},
		{
			name:      "identifier",
			tokenType: IDENTIFIER,
			want:      "<identifier>",
		},
		{
			name:      "unkown type",
			tokenType: TokenType(1 << 10),
			want:      "<unknown>",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tokenType.String()
			assert.Equalf(t, tt.want, got, "(TokenType).String() = %s, want %s", got, tt.want)
		})
	}
}

func TestToken_String(t *testing.T) {
	testCases := []struct {
		name  string
		token Token
		want  string
	}{
		{
			name: "bool",
			token: Token{
				Type:    BOOL,
				Literal: "BOOL",
			},
			want: "<bool> BOOL",
		},

		{
			name: "identifier",
			token: Token{
				Type:    IDENTIFIER,
				Literal: "hello",
			},
			want: "<identifier> hello",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.token.String()
			assert.Equalf(t, tt.want, got, "(Token).String() = %s, want %s", got, tt.want)
		})
	}
}
