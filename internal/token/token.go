package token

import (
	"fmt"
)

type TokenType int

const unknownType = "<unknown>"

const (
	LPAREN    TokenType = iota // (
	RPAREN                     // )
	LBRACE                     // {
	RBRACE                     // }
	COMMA                      // ,
	DOT                        // .
	PLUS                       // +
	MINUS                      // -
	SLASH                      // /
	STAR                       // *
	ASIGN                      // =
	SEMICOLON                  // ;

	BOOL
	TRUE
	FALSE
	FLOAT
	STRING
	INTEGER
	IDENTIFIER

	FN
	OR
	AND
	FOR
	IF
	VAR
	RETUNR
	INT

	EOF

	BAD_SYNTAX
)

var keywords = map[string]TokenType{
	"fn":     FN,
	"or":     OR,
	"and":    AND,
	"for":    FOR,
	"if":     IF,
	"var":    VAR,
	"return": RETUNR,
	"int":    INT,
}

func IsKeyword(s string) (TokenType, bool) {
	if t, ok := keywords[s]; ok {
		return t, true
	}

	return IDENTIFIER, false
}

var typeMap = map[TokenType]string{
	BAD_SYNTAX: "bad syntax",

	LPAREN:    "(",
	RPAREN:    ")",
	LBRACE:    "{",
	RBRACE:    "}",
	COMMA:     ",",
	DOT:       ".",
	PLUS:      "+",
	MINUS:     "-",
	SLASH:     "/",
	STAR:      "*",
	ASIGN:     "=",
	SEMICOLON: ";",

	BOOL:       "<bool>",
	FLOAT:      "<float>",
	STRING:     "<string>",
	INTEGER:    "<integer>",
	IDENTIFIER: "<identifier>",

	FN:     "fn",
	OR:     "or",
	AND:    "and",
	FOR:    "for",
	IF:     "if",
	VAR:    "var",
	RETUNR: "return",
	INT:    "int",

	EOF: "end of file",
}

func (t TokenType) String() string {
	if s, ok := typeMap[t]; ok {
		return s
	}

	return unknownType
}

func (t TokenType) EOF() bool {
	return t == EOF
}

func (t TokenType) Invalid() bool {
	return t == BAD_SYNTAX
}

func (t TokenType) Is(other TokenType) bool {
	return t == other
}

type Token struct {
	Type    TokenType
	Literal string
}

func New(tokenType TokenType) Token {
	return Token{
		Type: tokenType,
	}
}

func WithLiteral(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s", t.Type.String(), t.Literal)
}
