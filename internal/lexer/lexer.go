package lexer

import (
	"errors"
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/Pshimaf-Git/corvus/internal/token"
)

var (
	ErrInvalidCharacter error = errors.New("invalid character")
)

const (
	lparen     = '('
	rparen     = ')'
	lbrace     = '{'
	rbrace     = '}'
	comma      = ','
	dot        = '.'
	plus       = '+'
	minus      = '-'
	slash      = '/'
	star       = '*'
	asign      = '='
	semicolon  = ';'
	emptyBlank = '_'
)

type Lexer struct {
	source string
	pos    int
}

func New(code string) Lexer {
	return Lexer{
		source: code,
		pos:    0,
	}
}

func (l *Lexer) Source() string {
	return l.source
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	if l.pos >= len(l.source) {
		return token.New(token.EOF)
	}

	cher, size := utf8.DecodeRuneInString(l.source[l.pos:])

	switch {
	case unicode.IsLetter(cher), cher == emptyBlank:
		return l.readIdentifier()
	case unicode.IsDigit(cher):
		return l.readNumder()
	case cher == '"':
		return l.readString()
	default:
		switch cher {
		case asign:
			l.pos += size
			return token.WithLiteral(token.ASIGN, string(asign))

		case plus:
			l.pos += size
			return token.WithLiteral(token.PLUS, string(plus))

		case minus:
			l.pos += size
			return token.WithLiteral(token.MINUS, string(minus))

		case slash:
			l.pos += size
			return token.WithLiteral(token.SLASH, string(slash))

		case star:
			l.pos += size
			return token.WithLiteral(token.STAR, string(star))

		case lparen:
			l.pos += size
			return token.WithLiteral(token.LPAREN, string(lparen))

		case rparen:
			l.pos += size
			return token.WithLiteral(token.RPAREN, string(rparen))

		case lbrace:
			l.pos += size
			return token.WithLiteral(token.LBRACE, string(lbrace))

		case rbrace:
			l.pos += size
			return token.WithLiteral(token.RBRACE, string(rbrace))

		case comma:
			l.pos += size
			return token.WithLiteral(token.COMMA, string(comma))

		case dot:
			l.pos += size
			return token.WithLiteral(token.DOT, string(dot))

		case semicolon:
			l.pos += size
			return token.WithLiteral(token.SEMICOLON, string(semicolon))
		}
	}

	l.pos += size
	return token.WithLiteral(token.BAD_SYNTAX, fmt.Sprintf("unknown %s", string(cher)))
}

func (l *Lexer) Process() ([]token.Token, error) {
	var t token.Token

	res := []token.Token(nil)

	for !t.Type.EOF() {
		t = l.NextToken()
		if t.Type.Invalid() {
			return nil, fmt.Errorf("%w: %s", ErrInvalidCharacter, t.String())
		}

		res = append(res, t)
	}

	return res, nil
}

func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.source) {
		r, size := utf8.DecodeRuneInString(l.source[l.pos:])
		if unicode.IsSpace(r) {
			l.pos += size
		} else {
			break
		}
	}
}

func (l *Lexer) readIdentifier() token.Token {
	start := l.pos

	for l.pos < len(l.source) {
		r, size := utf8.DecodeRuneInString(l.source[l.pos:])
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == emptyBlank {
			l.pos += size
		} else {
			break
		}
	}

	literal := l.source[start:l.pos]
	tokenType, ok := token.IsKeyword(literal)
	if ok {
		return token.WithLiteral(tokenType, literal)
	}

	return token.WithLiteral(token.IDENTIFIER, literal)
}

func (l *Lexer) readNumder() token.Token {
	start := l.pos

	for l.pos < len(l.source) {
		r, size := utf8.DecodeRuneInString(l.source[l.pos:])
		if unicode.IsDigit(r) {
			l.pos += size
		} else {
			break
		}
	}

	return token.WithLiteral(token.INTEGER, l.source[start:l.pos])
}

func (l *Lexer) readString() token.Token {
	// skip "
	l.pos++

	start := l.pos

	for l.pos < len(l.source) {
		r, size := utf8.DecodeRuneInString(l.source[l.pos:])
		if r == '"' {
			literal := l.source[start:l.pos]
			l.pos++
			return token.WithLiteral(token.STRING, literal)
		}
		l.pos += size
	}

	return token.WithLiteral(token.BAD_SYNTAX, "expactdstring literal must be closed")
}
